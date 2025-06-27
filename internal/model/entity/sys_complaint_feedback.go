// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComplaintFeedback is the golang structure for table sys_complaint_feedback.
type SysComplaintFeedback struct {
	Id               int64       `json:"id"               orm:"id"                ` // 自增主键ID
	SurveyCode       string      `json:"surveyCode"       orm:"survey_code"       ` // 问卷编号
	TicketNo         string      `json:"ticketNo"         orm:"ticket_no"         ` // 投诉标识
	InvestigatorName string      `json:"investigatorName" orm:"investigator_name" ` // 调查者姓名
	ContactInfo      string      `json:"contactInfo"      orm:"contact_info"      ` // 联系方式 (电话或邮箱)
	ProcessingSpeed  string      `json:"processingSpeed"  orm:"processing_speed"  ` // 处理速度
	StaffAttitude    string      `json:"staffAttitude"    orm:"staff_attitude"    ` // 处理人员态度
	ResolutionEffect string      `json:"resolutionEffect" orm:"resolution_effect" ` // 问题解决效果
	OtherSuggestions string      `json:"otherSuggestions" orm:"other_suggestions" ` // 其它建议或意见
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        ` // 创建时间
	IsDeleted        int         `json:"isDeleted"        orm:"is_deleted"        ` // 逻辑删除标志 (0: 未删除, 1: 已删除)
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"        ` // 删除时间
}
