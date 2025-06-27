// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComplaints is the golang structure for table sys_complaints.
type SysComplaints struct {
	Id                 int         `json:"id"                 orm:"id"                  ` //
	TicketNo           int64       `json:"ticketNo"           orm:"ticket_no"           ` // 投诉标识
	Title              string      `json:"title"              orm:"title"               ` // 投诉标题
	Content            string      `json:"content"            orm:"content"             ` // 投诉内容
	TypeCode           string      `json:"typeCode"           orm:"type_code"           ` // 投诉类型编码
	SourceCode         string      `json:"sourceCode"         orm:"source_code"         ` // 投诉来源编码
	PriorityCode       string      `json:"priorityCode"       orm:"priority_code"       ` // 投诉等级编码
	StatusCode         string      `json:"statusCode"         orm:"status_code"         ` // 处理状态编码
	ComplainantName    string      `json:"complainantName"    orm:"complainant_name"    ` // 投诉人姓名
	ComplainantContact string      `json:"complainantContact" orm:"complainant_contact" ` // 投诉人联系方式
	AreaCode           string      `json:"areaCode"           orm:"area_code"           ` // 投诉区域编码
	AssigneeId         int         `json:"assigneeId"         orm:"assignee_id"         ` // 负责人ID
	Satisfaction       string      `json:"satisfaction"       orm:"satisfaction"        ` // 满意度评价
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          ` // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          ` // 更新时间
	CompletedAt        *gtime.Time `json:"completedAt"        orm:"completed_at"        ` // 完成时间
	FeedbackTime       *gtime.Time `json:"feedbackTime"       orm:"feedback_time"       ` // 反馈时间
	ProcessingNotes    string      `json:"processingNotes"    orm:"processing_notes"    ` // 处理备注
	IsDeleted          int         `json:"isDeleted"          orm:"is_deleted"          ` // 逻辑删除标志 (0: 未删除, 1: 已删除)
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"          ` // 删除时间
}
