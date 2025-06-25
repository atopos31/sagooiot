package system

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type SysFsListReq struct {
	g.Meta `path:"/file/list" tags:"文件管理" method:"get" summary:"获取文件列表"`
	Path   string `v:"required" dc:"路径"`
	*common.PaginationReq
}

type SysFsDoRes struct {
	g.Meta `mime:"application/json"`
	Data   []SysFsDoResData `json:"Data"`
	common.PaginationRes
}

type SysFsDoResData struct {
	Id       uint64      `json:"id"`
	Name     string      `json:"name"`
	Title    string      `json:"title"`
	Remark   string      `json:"remark"`
	Size     uint64      `json:"size"`
	IsDir    bool        `json:"isDir"`
	UpdateAt *gtime.Time `json:"updateAt"`
}

type SysFsDelReq struct {
	g.Meta `path:"/file/del" tags:"文件管理" method:"post" summary:"删除文件"`
	Id     uint64 `json:"id" v:"required" dc:"文件/目录ID"`
	IsDir  bool   `json:"isDir" dc:"是否是目录"`
}

type SysFsDelRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id"`
	IsDir  bool   `json:"isDir"`
}

type SysFsUploadReq struct {
	g.Meta `path:"/file/upload" tags:"文件管理" method:"post" summary:"上传文件"`
	Title  string            `json:"title"  dc:"文件标题"`
	Path   string            `json:"path" v:"required" dc:"上传路径"`
	Remark string            `json:"remark" dc:"备注"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type SysFsUploadRes struct {
	g.Meta `mime:"application/json"`
}

type SysCreateDirReq struct {
	g.Meta  `path:"/file/create/dir" tags:"文件管理" method:"post" summary:"创建目录"`
	Path    string `json:"path" v:"required" dc:"目录父路径"`
	Name    string `json:"name" v:"required" dc:"目录名称"`
	Remarks string `json:"remark" dc:"备注"`
}

type SysCreateDirRes struct {
	g.Meta `mime:"application/json"`
}

type SysFileDownloadReq struct {
	g.Meta `path:"/file/download" tags:"文件管理" method:"get" summary:"下载文件"`
	Id     uint64 `v:"required" dc:"文件ID"`
}

type SysFsDownloadRes struct {
	g.Meta `mime:"application/json"`
}

type SysFsSearchFileReq struct {
	g.Meta `path:"/file/search" tags:"文件管理" method:"get" summary:"搜索文件"`
	Query  string `json:"query" v:"required" dc:"搜索内容"`
	*common.PaginationReq
}

type SysFsDirTreeReq struct {
	g.Meta `path:"/file/dir/tree" tags:"文件管理" method:"get" summary:"目录树"`
}

type SysFsTreeRes struct {
	g.Meta `mime:"application/json"`
	Dirs   []*model.DirItemNode `json:"dirs"`
}
