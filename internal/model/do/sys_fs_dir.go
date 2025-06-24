// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFsDir is the golang structure of table sys_fs_dir for DAO operations like Where/Data.
type SysFsDir struct {
	g.Meta    `orm:"table:sys_fs_dir, do:true"`
	Id        interface{} // 主键ID
	ParentId  interface{} // 父目录ID，根目录此项为NULL
	Name      interface{} // 目录名称
	FullPath  interface{} // 完整路径（冗余字段，用于快速查询）
	Remarks   interface{} // 备注信息
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
	IsDeleted interface{} // 删除标志（0=未删除, 1=已删除）
	DeletedAt *gtime.Time // 软删除时间
	DeletedBy interface{} // 删除者ID
}
