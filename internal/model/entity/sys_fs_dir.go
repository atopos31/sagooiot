// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFsDir is the golang structure for table sys_fs_dir.
type SysFsDir struct {
	Id        uint64      `json:"id"        orm:"id"         ` // 主键ID
	ParentId  uint64      `json:"parentId"  orm:"parent_id"  ` // 父目录ID，根目录此项为NULL
	Name      string      `json:"name"      orm:"name"       ` // 目录名称
	FullPath  string      `json:"fullPath"  orm:"full_path"  ` // 完整路径（冗余字段，用于快速查询）
	Remarks   string      `json:"remarks"   orm:"remarks"    ` // 备注信息
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	IsDeleted int         `json:"isDeleted" orm:"is_deleted" ` // 删除标志（0=未删除, 1=已删除）
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 软删除时间
	DeletedBy uint64      `json:"deletedBy" orm:"deleted_by" ` // 删除者ID
}
