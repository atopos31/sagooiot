// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComplaintFeedback is the golang structure of table sys_complaint_feedback for DAO operations like Where/Data.
type SysComplaintFeedback struct {
	g.Meta           `orm:"table:sys_complaint_feedback, do:true"`
	Id               interface{} // 自增主键ID
	SurveyCode       interface{} // 问卷编号
	TicketNo         interface{} // 投诉标识
	InvestigatorName interface{} // 调查者姓名
	ContactInfo      interface{} // 联系方式 (电话或邮箱)
	ProcessingSpeed  interface{} // 处理速度
	StaffAttitude    interface{} // 处理人员态度
	ResolutionEffect interface{} // 问题解决效果
	OtherSuggestions interface{} // 其它建议或意见
	CreatedAt        *gtime.Time // 创建时间
	IsDeleted        interface{} // 逻辑删除标志 (0: 未删除, 1: 已删除)
	DeletedAt        *gtime.Time // 删除时间
}
