// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysFsFileDao is the data access object for the table sys_fs_file.
type SysFsFileDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysFsFileColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysFsFileColumns defines and stores column names for the table sys_fs_file.
type SysFsFileColumns struct {
	Id          string // 文件唯一ID
	Name        string // 文件名称（含扩展名）
	DirectoryId string // 所属目录的ID
	Size        string // 文件大小（字节）
	Title       string // 文件标题
	Remarks     string // 备注信息
	UpdatedAt   string // 更新时间
	CreatedAt   string // 创建时间
	IsDeleted   string // 删除标志（0=未删除, 1=已删除）
	DeletedAt   string // 软删除时间
	DeletedBy   string // 删除者ID
}

// sysFsFileColumns holds the columns for the table sys_fs_file.
var sysFsFileColumns = SysFsFileColumns{
	Id:          "id",
	Name:        "name",
	DirectoryId: "directory_id",
	Size:        "size",
	Title:       "title",
	Remarks:     "remarks",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
	IsDeleted:   "is_deleted",
	DeletedAt:   "deleted_at",
	DeletedBy:   "deleted_by",
}

// NewSysFsFileDao creates and returns a new DAO object for table data access.
func NewSysFsFileDao(handlers ...gdb.ModelHandler) *SysFsFileDao {
	return &SysFsFileDao{
		group:    "default",
		table:    "sys_fs_file",
		columns:  sysFsFileColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysFsFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysFsFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysFsFileDao) Columns() SysFsFileColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysFsFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysFsFileDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysFsFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
