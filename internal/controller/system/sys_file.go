package system

import (
	"context"
	systemV1 "sagooiot/api/v1/system"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var SysFs = cSysFs{}

type cSysFs struct{}

func (c *cSysFs) List(ctx context.Context, req *systemV1.SysFsListReq) (res *systemV1.SysFsDoRes, err error) {
	g.Log().Debugf(ctx, "req: %+v", req)
	total,files, err := service.SysFile().GetFileList(ctx, req.Path,req.PageNum,req.PageSize)
	if err != nil {
		return nil, err
	}
	res = new(systemV1.SysFsDoRes)
	res.Data = make([]systemV1.SysFsDoResData, len(files))
	for i, v := range files {
		res.Data[i] = systemV1.SysFsDoResData{
			Id:       v.ID,
			IsDir:    v.IsDir,
			Name:     v.Name,
			Title:    v.Title,
			Remark:   v.Remarks,
			Size:     v.Size,
			UpdateAt: v.UpdateAt,
		}
	}
	res.Total = total
	res.CurrentPage = req.PageNum
	return
}

func (s *cSysFs) CreateDir(ctx context.Context, req *systemV1.SysCreateDirReq) (res *systemV1.SysFsDoRes, err error) {
	g.Log().Debugf(ctx, "req: %+v", req)
	err = service.SysFile().CreateDir(ctx, req.Path, req.Remarks, req.Name)
	if err != nil {
		return
	}
	return nil, nil
}

func (s *cSysFs) UploadFile(ctx context.Context, req *systemV1.SysFsUploadReq) (res *systemV1.SysFsDelRes, err error) {
	var title string
	if req.Title != "" {
		title = req.Title
	} else {
		title = req.File.Filename
	}
	err = service.SysFile().UploadFile(ctx, req.Path, title, req.Remark, req.File)
	return nil, err
}

func (c *cSysFs) Del(ctx context.Context, req *systemV1.SysFsDelReq) (res *systemV1.SysFsDelRes, err error) {
	g.Log().Debugf(ctx, "req: %+v", req)
	err = service.SysFile().DelFile(ctx, req.Id, req.IsDir)
	return &systemV1.SysFsDelRes{Id: req.Id, IsDir: req.IsDir}, err
}

func (c *cSysFs) Download(ctx context.Context, req *systemV1.SysFileDownloadReq) (res *systemV1.SysFsDownloadRes, err error) {
	fullPath, err := service.SysFile().GetFullPath(ctx, req.Id)
	if err != nil {
		return
	}
	request := g.RequestFromCtx(ctx)
	g.Log().Debugf(ctx, "fullPath: %s", fullPath)
	err = service.FileSystem().Download(ctx, request.Response, fullPath)
	if err != nil {
		return
	}
	return
}

func (c *cSysFs) SearchFile(ctx context.Context, req *systemV1.SysFsSearchFileReq) (res *systemV1.SysFsDoRes, err error) {
	g.Log().Debugf(ctx, "SearchFile req: %+v", req)
	total,out, err := service.SysFile().SearchFile(ctx, req.Query,req.PageNum,req.PageSize)
	res = new(systemV1.SysFsDoRes)
	res.Data = make([]systemV1.SysFsDoResData, len(out))
	for i, v := range out {
		res.Data[i] = systemV1.SysFsDoResData{
			Id:       v.ID,
			IsDir:    v.IsDir,
			Name:     v.Name,
			Title:    v.Title,
			Remark:   v.Remarks,
			Size:     v.Size,
			UpdateAt: v.UpdateAt,
		}
	}
	res.Total = total
	res.CurrentPage = req.PageNum
	return 
}

func (c *cSysFs) DirTre(ctx context.Context, req *systemV1.SysFsDirTreeReq) (res *systemV1.SysFsTreeRes, err error) {
	out,err  := service.SysFile().DirTree(ctx)
	if err != nil {
		return nil,err
	}
	res = new(systemV1.SysFsTreeRes)
	res.Dirs = out
	return
}