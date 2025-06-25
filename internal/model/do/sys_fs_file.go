// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFsFile is the golang structure of table sys_fs_file for DAO operations like Where/Data.
type SysFsFile struct {
	g.Meta      `orm:"table:sys_fs_file, do:true"`
	Id          interface{} // 文件唯一ID
	Name        interface{} // 文件名称（含扩展名）
	DirectoryId interface{} // 所属目录的ID
	Size        interface{} // 文件大小（字节）
	Title       interface{} // 文件标题
	Remarks     interface{} // 备注信息
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
	IsDeleted   interface{} // 删除标志（0=未删除, 1=已删除）
	DeletedAt   *gtime.Time // 软删除时间
	DeletedBy   interface{} // 删除者ID
}
