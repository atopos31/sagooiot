package model

import (

	"github.com/gogf/gf/v2/os/gtime"
)

type FileItem struct {
	ID          uint64      `orm:"id,primary" json:"id"`
	DirectoryId uint64      `orm:"directory_id" json:"directoryId"`
	Name        string      `orm:"name" json:"name"`
	IsDir       bool        `orm:"is_dir" json:"isDir"`
	Size        uint64      `orm:"size" json:"size"`
	UpdateAt    *gtime.Time `orm:"updated_at" json:"updatedAt"`
}
