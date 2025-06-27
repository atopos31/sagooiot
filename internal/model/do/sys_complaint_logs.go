// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComplaintLogs is the golang structure of table sys_complaint_logs for DAO operations like Where/Data.
type SysComplaintLogs struct {
	g.Meta        `orm:"table:sys_complaint_logs, do:true"`
	Id            interface{} //
	ComplaintId   interface{} // 投诉编号
	ActionType    interface{} // 操作类型
	ActionUser    interface{} // 操作人
	ActionContent interface{} // 操作内容
	CreatedAt     *gtime.Time // 操作时间
}
