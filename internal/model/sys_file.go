package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type FileItem struct {
	ID          uint64      `orm:"id,primary" json:"id"`
	DirectoryId uint64      `orm:"directory_id" json:"directoryId"`
	Title       string      `orm:"title" json:"title"`
	Remarks     string      `orm:"remarks" json:"remarks"`
	Name        string      `orm:"name" json:"name"`
	IsDir       bool        `orm:"is_dir" json:"isDir"`
	Size        uint64      `orm:"size" json:"size"`
	UpdateAt    *gtime.Time `orm:"updated_at" json:"updatedAt"`
}

type DirItemNode struct {
	Id        uint64        `orm:"id,primary" json:"id"`
	Name      string        `orm:"name" json:"name"`
	ParentId  *uint64       `orm:"parent_id"  json:"parentId"` // 父目录ID，根目录此项为NULL
	FullPath  string        `orm:"full_path" json:"fullPath"`
	Childrens []*DirItemNode `json:"childrens"`
}
