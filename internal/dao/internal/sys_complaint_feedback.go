// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysComplaintFeedbackDao is the data access object for the table sys_complaint_feedback.
type SysComplaintFeedbackDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  SysComplaintFeedbackColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// SysComplaintFeedbackColumns defines and stores column names for the table sys_complaint_feedback.
type SysComplaintFeedbackColumns struct {
	Id               string // 自增主键ID
	SurveyCode       string // 问卷编号
	TicketNo         string // 投诉标识
	InvestigatorName string // 调查者姓名
	ContactInfo      string // 联系方式 (电话或邮箱)
	ProcessingSpeed  string // 处理速度
	StaffAttitude    string // 处理人员态度
	ResolutionEffect string // 问题解决效果
	OtherSuggestions string // 其它建议或意见
	CreatedAt        string // 创建时间
	IsDeleted        string // 逻辑删除标志 (0: 未删除, 1: 已删除)
	DeletedAt        string // 删除时间
}

// sysComplaintFeedbackColumns holds the columns for the table sys_complaint_feedback.
var sysComplaintFeedbackColumns = SysComplaintFeedbackColumns{
	Id:               "id",
	SurveyCode:       "survey_code",
	TicketNo:         "ticket_no",
	InvestigatorName: "investigator_name",
	ContactInfo:      "contact_info",
	ProcessingSpeed:  "processing_speed",
	StaffAttitude:    "staff_attitude",
	ResolutionEffect: "resolution_effect",
	OtherSuggestions: "other_suggestions",
	CreatedAt:        "created_at",
	IsDeleted:        "is_deleted",
	DeletedAt:        "deleted_at",
}

// NewSysComplaintFeedbackDao creates and returns a new DAO object for table data access.
func NewSysComplaintFeedbackDao(handlers ...gdb.ModelHandler) *SysComplaintFeedbackDao {
	return &SysComplaintFeedbackDao{
		group:    "default",
		table:    "sys_complaint_feedback",
		columns:  sysComplaintFeedbackColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysComplaintFeedbackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysComplaintFeedbackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysComplaintFeedbackDao) Columns() SysComplaintFeedbackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysComplaintFeedbackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysComplaintFeedbackDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysComplaintFeedbackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
