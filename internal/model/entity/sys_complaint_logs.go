// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComplaintLogs is the golang structure for table sys_complaint_logs.
type SysComplaintLogs struct {
	Id            int         `json:"id"            orm:"id"             ` //
	ComplaintId   string      `json:"complaintId"   orm:"complaint_id"   ` // 投诉编号
	ActionType    string      `json:"actionType"    orm:"action_type"    ` // 操作类型
	ActionUser    string      `json:"actionUser"    orm:"action_user"    ` // 操作人
	ActionContent string      `json:"actionContent" orm:"action_content" ` // 操作内容
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     ` // 操作时间
}
