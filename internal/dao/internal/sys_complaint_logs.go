// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysComplaintLogsDao is the data access object for the table sys_complaint_logs.
type SysComplaintLogsDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  SysComplaintLogsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// SysComplaintLogsColumns defines and stores column names for the table sys_complaint_logs.
type SysComplaintLogsColumns struct {
	Id            string //
	ComplaintId   string // 投诉编号
	ActionType    string // 操作类型
	ActionUser    string // 操作人
	ActionContent string // 操作内容
	CreatedAt     string // 操作时间
}

// sysComplaintLogsColumns holds the columns for the table sys_complaint_logs.
var sysComplaintLogsColumns = SysComplaintLogsColumns{
	Id:            "id",
	ComplaintId:   "complaint_id",
	ActionType:    "action_type",
	ActionUser:    "action_user",
	ActionContent: "action_content",
	CreatedAt:     "created_at",
}

// NewSysComplaintLogsDao creates and returns a new DAO object for table data access.
func NewSysComplaintLogsDao(handlers ...gdb.ModelHandler) *SysComplaintLogsDao {
	return &SysComplaintLogsDao{
		group:    "default",
		table:    "sys_complaint_logs",
		columns:  sysComplaintLogsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysComplaintLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysComplaintLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysComplaintLogsDao) Columns() SysComplaintLogsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysComplaintLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysComplaintLogsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysComplaintLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
