// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFsFile is the golang structure for table sys_fs_file.
type SysFsFile struct {
	Id          uint64      `json:"id"          orm:"id"           ` // 文件唯一ID
	Name        string      `json:"name"        orm:"name"         ` // 文件名称（含扩展名）
	DirectoryId uint64      `json:"directoryId" orm:"directory_id" ` // 所属目录的ID
	Size        uint64      `json:"size"        orm:"size"         ` // 文件大小（字节）
	Title       string      `json:"title"       orm:"title"        ` // 文件标题
	Remarks     string      `json:"remarks"     orm:"remarks"      ` // 备注信息
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	IsDeleted   int         `json:"isDeleted"   orm:"is_deleted"   ` // 删除标志（0=未删除, 1=已删除）
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   ` // 软删除时间
	DeletedBy   uint64      `json:"deletedBy"   orm:"deleted_by"   ` // 删除者ID
}
