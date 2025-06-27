// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysComplaintsDao is the data access object for the table sys_complaints.
type SysComplaintsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysComplaintsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysComplaintsColumns defines and stores column names for the table sys_complaints.
type SysComplaintsColumns struct {
	Id                 string //
	TicketNo           string // 投诉标识
	Title              string // 投诉标题
	Content            string // 投诉内容
	TypeCode           string // 投诉类型编码
	SourceCode         string // 投诉来源编码
	PriorityCode       string // 投诉等级编码
	StatusCode         string // 处理状态编码
	ComplainantName    string // 投诉人姓名
	ComplainantContact string // 投诉人联系方式
	AreaCode           string // 投诉区域编码
	AssigneeId         string // 负责人ID
	Satisfaction       string // 满意度评价
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
	CompletedAt        string // 完成时间
	FeedbackTime       string // 反馈时间
	ProcessingNotes    string // 处理备注
	IsDeleted          string // 逻辑删除标志 (0: 未删除, 1: 已删除)
	DeletedAt          string // 删除时间
}

// sysComplaintsColumns holds the columns for the table sys_complaints.
var sysComplaintsColumns = SysComplaintsColumns{
	Id:                 "id",
	TicketNo:           "ticket_no",
	Title:              "title",
	Content:            "content",
	TypeCode:           "type_code",
	SourceCode:         "source_code",
	PriorityCode:       "priority_code",
	StatusCode:         "status_code",
	ComplainantName:    "complainant_name",
	ComplainantContact: "complainant_contact",
	AreaCode:           "area_code",
	AssigneeId:         "assignee_id",
	Satisfaction:       "satisfaction",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	CompletedAt:        "completed_at",
	FeedbackTime:       "feedback_time",
	ProcessingNotes:    "processing_notes",
	IsDeleted:          "is_deleted",
	DeletedAt:          "deleted_at",
}

// NewSysComplaintsDao creates and returns a new DAO object for table data access.
func NewSysComplaintsDao(handlers ...gdb.ModelHandler) *SysComplaintsDao {
	return &SysComplaintsDao{
		group:    "default",
		table:    "sys_complaints",
		columns:  sysComplaintsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysComplaintsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysComplaintsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysComplaintsDao) Columns() SysComplaintsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysComplaintsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysComplaintsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysComplaintsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
