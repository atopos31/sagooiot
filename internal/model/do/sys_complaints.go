// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComplaints is the golang structure of table sys_complaints for DAO operations like Where/Data.
type SysComplaints struct {
	g.Meta             `orm:"table:sys_complaints, do:true"`
	Id                 interface{} //
	TicketNo           interface{} // 投诉标识
	Title              interface{} // 投诉标题
	Content            interface{} // 投诉内容
	TypeCode           interface{} // 投诉类型编码
	SourceCode         interface{} // 投诉来源编码
	PriorityCode       interface{} // 投诉等级编码
	StatusCode         interface{} // 处理状态编码
	ComplainantName    interface{} // 投诉人姓名
	ComplainantContact interface{} // 投诉人联系方式
	AreaCode           interface{} // 投诉区域编码
	AssigneeId         interface{} // 负责人ID
	Satisfaction       interface{} // 满意度评价
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
	CompletedAt        *gtime.Time // 完成时间
	FeedbackTime       *gtime.Time // 反馈时间
	ProcessingNotes    interface{} // 处理备注
	IsDeleted          interface{} // 逻辑删除标志 (0: 未删除, 1: 已删除)
	DeletedAt          *gtime.Time // 删除时间
}
