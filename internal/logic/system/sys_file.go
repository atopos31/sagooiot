package system

import (
	"context"
	"errors"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
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

func (s *sSysFile) GetFileList(ctx context.Context, path string) ([]model.FileItem, error) {
	res, err := dao.SysFsDir.Ctx(ctx).
		Where(dao.SysFsDir.Columns().FullPath, path).
		Fields(dao.SysFsDir.Columns().Id).One()
	if err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, "res:", res)
	id := res.GMap().GetVar(dao.SysFsDir.Columns().Id).Val()
	dirs := []entity.SysFsDir{}
	if err := dao.SysFsDir.Ctx(ctx).
		Where(dao.SysFsDir.Columns().ParentId, id).
		Where(dao.SysFsDir.Columns().IsDeleted, 0).
		Scan(&dirs); err != nil {
		return nil, errors.New("获取目录列表失败")
	}
	g.Log().Debug(ctx, "dirs:", dirs)
	files := []entity.SysFsFile{}
	if err := dao.SysFsFile.Ctx(ctx).
		Where(dao.SysFsFile.Columns().DirectoryId, id).
		Where(dao.SysFsFile.Columns().IsDeleted, 0).
		Scan(&files); err != nil {
		return nil, errors.New("获取文件列表失败")
	}
	out := make([]model.FileItem, 0)
	for _, v := range dirs {
		out = append(out, model.FileItem{
			ID:       v.Id,
			IsDir:    true,
			Name:     v.Name,
			Size:     0,
			UpdateAt: v.UpdatedAt,
		})
	}
	for _, v := range files {
		out = append(out, model.FileItem{
			ID:       v.Id,
			IsDir:    false,
			Name:     v.Name,
			Size:     v.Size,
			UpdateAt: v.UpdatedAt,
		})
	}
	g.Log().Debug(ctx, "Run TaskDeviceDataSaveWorker: %v", out)

	return out, nil
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

	// 物理文件夹位置
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.path").String()
	if err := gfile.Mkdir(uploadPath + path + "/" + dir); err != nil {
		return err
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

func (s *sSysFile) UploadFile(ctx context.Context, path string, Remarks string, file *ghttp.UploadFile) (err error) {
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
	// 物理保存 TODO 对接多个文件系统
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.path").String()
	g.Log().Debug(ctx, "uploadPath:", uploadPath)
	_, err = file.Save(uploadPath+path, false)
	if err != nil {
		return gerror.New("文件保存到本地失败")
	}

	_, err = dao.SysFsFile.Ctx(ctx).Data(do.SysFsFile{
		Name:        file.Filename,
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

		// 物理删除
		uploadPath := g.Cfg().MustGet(ctx, "system.upload.path").String()
		path, err := dao.SysFsDir.Ctx(ctx).Where(dao.SysFsDir.Columns().Id, id).Value(dao.SysFsDir.Columns().FullPath)
		if err := gfile.Remove(uploadPath + path.String()); err != nil {
			return err
		}

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
		if err != nil {
			return err
		}
		if err := gfile.Remove(fullPath); err != nil {
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

// 获取文件的物理地址
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
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.path").String()
	return uploadPath + dir.String() + "/" + file.Name, nil
}
