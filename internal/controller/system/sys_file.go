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
	files, err := service.SysFile().GetFileList(ctx, req.Path)
	if err != nil {
		return nil, err
	}
	res = new(systemV1.SysFsDoRes)
	res.Files = make([]systemV1.SysFsDoResData, len(files))
	for i, v := range files {
		res.Files[i] = systemV1.SysFsDoResData{
			Id:       v.ID,
			IsDir:    v.IsDir,
			Name:     v.Name,
			Size:     v.Size,
			UpdateAt: v.UpdateAt,
		}
	}
	return
}

func (s *cSysFs) CreateDir(ctx context.Context, req *systemV1.SysCreateDirReq) (res *systemV1.SysFsDoRes, err error) {
	g.Log().Debugf(ctx, "req: %+v", req)
	err = service.SysFile().CreateDir(ctx, req.Path, req.Remarks, req.Name)
	if err != nil {
		return
	}
	return nil,nil
}

func (s *cSysFs) UploadFile(ctx context.Context, req *systemV1.SysFsUploadReq) (res *systemV1.SysFsDelRes, err error) {
	err = service.SysFile().UploadFile(ctx, req.Path, req.Remarks, req.File)
	return nil,err
}

func (c *cSysFs) Del(ctx context.Context, req *systemV1.SysFsDelReq) (res *systemV1.SysFsDelRes, err error) {
	g.Log().Debugf(ctx, "req: %+v", req)
	err = service.SysFile().DelFile(ctx, req.Id, req.IsDir)
	return &systemV1.SysFsDelRes{Id: req.Id, IsDir: req.IsDir}, err
}

func (c *cSysFs) Download(ctx context.Context, req *systemV1.SysFileDownloadReq) (res *systemV1.SysFsDownloadRes, err error) {
	fullPath,err := service.SysFile().GetFullPath(ctx, req.Id)
	if err != nil {
		return
	}
	request := g.RequestFromCtx(ctx)
	g.Log().Debugf(ctx, "fullPath: %s", fullPath)
	request.Response.ServeFileDownload(fullPath)
	return
}
