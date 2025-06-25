package system

import (
	"context"
	"errors"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
	"slices"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type sSysFile struct {
}

func sysFileNew() *sSysFile {
	return &sSysFile{}
}

func init() {
	service.RegisterSysFile(sysFileNew())
}

func (s *sSysFile) GetFileList(ctx context.Context, path string, num int, size int) (int, []model.FileItem, error) {
	res, err := dao.SysFsDir.Ctx(ctx).
		Where(dao.SysFsDir.Columns().FullPath, path).
		Fields(dao.SysFsDir.Columns().Id).One()
	if err != nil {
		return 0, nil, err
	}
	g.Log().Debug(ctx, "res:", res)
	id := res.GMap().GetVar(dao.SysFsDir.Columns().Id).Val()

	m := dao.SysFsFile.Ctx(ctx).Where(dao.SysFsFile.Columns().DirectoryId, id).Where(dao.SysFsFile.Columns().IsDeleted, 0)
	total, err := m.Count()
	if err != nil {
		return 0, nil, err
	}
	files := []model.FileItem{}
	if err := m.Page(num, size).OrderDesc(dao.SysFsFile.Columns().CreatedAt).Scan(&files); err != nil {
		return 0, nil, err
	}

	return total, files, nil
}

func (s *sSysFile) CreateDir(ctx context.Context, path string, Remarks string, dir string) (err error) {
	res, err := dao.SysFsDir.Ctx(ctx).
		Where(dao.SysFsDir.Columns().FullPath, path).
		Fields(dao.SysFsDir.Columns().Id).One()
	if err != nil {
		return err
	}
	g.Log().Debug(ctx, "create dir res:", res)
	id := res.GMap().GetVar(dao.SysFsDir.Columns().Id).Val()

	if path == "/" {
		path = ""
	}
	// FullPath 是否存在
	count, err := dao.SysFsDir.Ctx(ctx).Where(dao.SysFsDir.Columns().FullPath, path+"/"+dir).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("目录已存在")
	}

	_, err = dao.SysFsDir.Ctx(ctx).Data(do.SysFsDir{
		Name:      dir,
		FullPath:  path + "/" + dir,
		Remarks:   Remarks,
		UpdatedAt: gtime.Now(),
		ParentId:  id,
		CreatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysFile) UploadFile(ctx context.Context, path string, title string, Remarks string, file *ghttp.UploadFile) (err error) {
	res, err := dao.SysFsDir.Ctx(ctx).
		Where(dao.SysFsDir.Columns().FullPath, path).
		Fields(dao.SysFsDir.Columns().Id).One()
	if err != nil {
		return err
	}
	if path != "/" && res == nil {
		return gerror.New("目标目录不存在")
	}
	g.Log().Debug(ctx, "res:", res)
	id := res.GMap().GetVar(dao.SysFsDir.Columns().Id).Val()
	// 检查该目录是否存在此文件
	count, err := dao.SysFsFile.Ctx(ctx).
		Where(dao.SysFsFile.Columns().DirectoryId, id).
		Where(dao.SysFsFile.Columns().Name, file.Filename).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("文件已存在，无法重复创建")
	}
	// 物理保存
	if err := service.FileSystem().Save(ctx, file, path); err != nil {
		return err
	}

	_, err = dao.SysFsFile.Ctx(ctx).Data(do.SysFsFile{
		Name:        file.Filename,
		Title:       title,
		Remarks:     Remarks,
		Size:        file.Size,
		UpdatedAt:   gtime.Now(),
		DirectoryId: id,
		CreatedAt:   gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysFile) DelFile(ctx context.Context, id uint64, isDir bool) (err error) {
	userId := service.Context().GetUserId(ctx)
	if isDir {
		_, err = dao.SysFsDir.Ctx(ctx).Where(dao.SysFsDir.Columns().Id, id).Data(g.Map{
			dao.SysFsDir.Columns().DeletedAt: gtime.Now(),
			dao.SysFsDir.Columns().IsDeleted: 1,
			dao.SysFsDir.Columns().DeletedBy: userId,
		}).Update()
		if err != nil {
			return err
		}
	} else {
		// 物理删除
		fullPath, err := s.GetFullPath(ctx, id)
		if err := service.FileSystem().Remove(ctx, fullPath); err != nil {
			return err
		}

		_, err = dao.SysFsFile.Ctx(ctx).Where(dao.SysFsFile.Columns().Id, id).Data(g.Map{
			dao.SysFsFile.Columns().DeletedAt: gtime.Now(),
			dao.SysFsFile.Columns().IsDeleted: 1,
			dao.SysFsFile.Columns().DeletedBy: userId,
		}).Update()
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取文件的完整地址
func (s *sSysFile) GetFullPath(ctx context.Context, id uint64) (fullPath string, err error) {
	file := model.FileItem{}
	err = dao.SysFsFile.Ctx(ctx).Where(dao.SysFsFile.Columns().Id, id).Scan(&file)
	if err != nil {
		return
	}
	dir, err := dao.SysFsDir.Ctx(ctx).Where(dao.SysFsDir.Columns().Id, file.DirectoryId).Value(dao.SysFsDir.Columns().FullPath)
	if err != nil {
		return
	}
	return dir.String() + "/" + file.Name, nil
}

func (s *sSysFile) SearchFile(ctx context.Context, query string, num int, size int) (int, []model.FileItem, error) {
	files := []model.FileItem{}
	m := dao.SysFsFile.Ctx(ctx).
		WhereLike(dao.SysFsFile.Columns().Name, "%"+query+"%").
		WhereOrLike(dao.SysFsFile.Columns().Title, "%"+query+"%").
		WhereOrLike(dao.SysFsFile.Columns().Remarks, "%"+query+"%")
	total, err := m.Count()
	if err != nil {
		return 0, nil, errors.New("获取文件列表失败")
	}
	err = m.Fields(model.FileItem{}).Page(num, size).
		OrderDesc(dao.SysFsFile.Columns().CreatedAt).Scan(&files)
	if err != nil {
		return 0, nil, errors.New("获取文件列表失败")
	}
	return total, files, nil
}

func (s *sSysFile) DirTree(ctx context.Context) ([]*model.DirItemNode, error) {
	var dirs []model.DirItemNode
	err := dao.SysFsDir.Ctx(ctx).
		Where(dao.SysFsDir.Columns().IsDeleted, 0).
		Scan(&dirs)
	if err != nil {
		return nil, err
	}
	g.Log().Debugf(ctx, "dirs: %+v", dirs)

	nodes := make(map[uint64]*model.DirItemNode, 0)
	for _, v := range dirs {
		nodes[v.Id] = &model.DirItemNode{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			FullPath:  v.FullPath,
			Childrens: make([]*model.DirItemNode, 0),
		}
	}
	g.Log().Debugf(ctx, "nodes: %+v", nodes)
	//根目录文件夹
	var roots []*model.DirItemNode
	for _, node := range nodes {
		if node.ParentId == nil {
			roots = append(roots, node)
			continue
		}

		//存在父节点 将其加入到父节点的子节点
		if parent, ok := nodes[*node.ParentId]; ok {
			parent.Childrens = append(parent.Childrens, node)
			slices.SortFunc(parent.Childrens, func(a, b *model.DirItemNode) int {
				return strings.Compare(a.Name, b.Name)
			})
			continue
		}
	}
	slices.SortFunc(roots,func(a ,b *model.DirItemNode) int {
		return strings.Compare(a.Name, b.Name)
	})

	return roots, nil
}
