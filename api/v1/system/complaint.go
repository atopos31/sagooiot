package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// AddComplaintReq 添加投诉
type AddComplaintReq struct {
	g.Meta          `path:"/complaint/add" tags:"投诉管理" method:"post" summary:"添加投诉"`
	Title           string `json:"title" v:"required#投诉标题不能为空" dc:"投诉标题"`
	Category        string `json:"category" v:"required#投诉类型不能为空" dc:"投诉类型(数据字典维护)"`
	Source          string `json:"source" v:"required#投诉来源不能为空" dc:"投诉来源（数据字典维护）"`
	Area            string `json:"area" v:"required#投诉区域不能为空" dc:"投诉区域(数据字典维护)"`
	ComplainantName string `json:"complainantName" v:"required#投诉人姓名不能为空" dc:"投诉人姓名"`
	Contact         string `json:"contact" dc:"联系方式"`
	Level           string `json:"level" v:"required#投诉等级不能为空" dc:"投诉等级(数据字典维护)"`
	Content         string `json:"content" v:"required#投诉内容不能为空" dc:"投诉内容"`
	Assignee        int    `json:"assignee" dc:"指派负责人"`
}

// AddComplaintRes 添加投诉
type AddComplaintRes struct {
}

// ComplaintListReq 投诉列表
type ComplaintListReq struct {
	g.Meta `path:"/complaint/list" tags:"投诉管理" method:"get" summary:"投诉列表"`
	common.PaginationReq
	Name     string `json:"name" dc:"名称"`
	Status   string `json:"status" dc:"状态(数据字典维护)"`
	Category string `json:"category" dc:"分类(数据字典维护)"`
	Level    string `json:"level" dc:"等级(数据字典维护)"`
}

// ComplaintListRes 投诉列表
type ComplaintListRes struct {
	common.PaginationRes
	Data []*model.ComplaintInfoRes `json:"Data" dc:"投诉列表数据"`
}

// GetComplaintByIdReq 根据TicketNo获取投诉信息
type GetComplaintByIdReq struct {
	g.Meta `path:"/complaint/info" tags:"投诉管理" method:"get" summary:"根据TicketNo获取投诉信息"`
	Id     int64 `json:"id" v:"required#TicketNo不能为空" dc:"投诉单号"`
}

// GetComplaintByIdRes 根据ID获取投诉信息
type GetComplaintByIdRes struct {
	Data *model.ComplaintInfoRes `json:"Data" dc:"投诉详情数据"`
}

// EditComplaintReq 投诉编辑
type EditComplaintReq struct {
	g.Meta          `path:"/complaint/edit" tags:"投诉管理" method:"put" summary:"投诉编辑"`
	Id              int64  `json:"id" v:"required#投诉单号不能为空" dc:"投诉单号"`
	Title           string `json:"title" v:"required#投诉标题不能为空" dc:"投诉标题"`
	Category        string `json:"category" v:"required#投诉类别不能为空" dc:"投诉类别"`
	Source          string `json:"source" v:"required#投诉来源不能为空" dc:"投诉来源"`
	Area            string `json:"area" v:"required#投诉区域不能为空" dc:"投诉区域"`
	ComplainantName string `json:"complainantName" v:"required#投诉人姓名不能为空" dc:"投诉人姓名"`
	Contact         string `json:"contact" dc:"联系方式"`
	Level           string `json:"level" v:"required#投诉等级不能为空" dc:"投诉等级"`
	Content         string `json:"content" v:"required#投诉内容不能为空" dc:"投诉内容"`
	Assignee        int    `json:"assignee" dc:"指派负责人"`
}

// EditComplaintRes 投诉编辑
type EditComplaintRes struct {
}

// DeleteComplaintByIdReq 批量删除投诉
type DeleteComplaintByIdReq struct {
	g.Meta `path:"/complaint/delete" tags:"投诉管理" method:"delete" summary:"批量删除投诉"`
	Ids    []int64 `json:"ids" v:"required#ID列表不能为空" dc:"投诉单号列表"`
}

// DeleteComplaintByIdRes 批量删除投诉
type DeleteComplaintByIdRes struct {
}

// ComplaintOverviewReq 获取投诉概要统计
type ComplaintOverviewReq struct {
	g.Meta    `path:"/complaint/overview" tags:"投诉管理" method:"get" summary:"获取投诉概要统计"`
	TimeRange string `json:"timeRange" dc:"时间范围: week/month/quarter/year"`
}

// ComplaintOverviewRes 投诉概要统计响应
type ComplaintOverviewRes struct {
	Data *ComplaintOverviewData `json:"Data" dc:"统计数据"`
}

// ComplaintOverviewData 投诉概要统计数据
type ComplaintOverviewData struct {
	TotalComplaints       int     `json:"totalComplaints" dc:"总投诉数"`
	PendingComplaints     int     `json:"pendingComplaints" dc:"待处理投诉数"`
	CompletedComplaints   int     `json:"completedComplaints" dc:"已完成投诉数"`
	UrgentComplaints      int     `json:"urgentComplaints" dc:"紧急投诉数"`
	AverageProcessingTime float64 `json:"averageProcessingTime" dc:"平均处理时间(天)"`
	CompletionRate        float64 `json:"completionRate" dc:"完成率(%)"`
	SatisfactionScore     float64 `json:"satisfactionScore" dc:"满意度评分"`
	SatisfactionTotal     int     `json:"satisfactionTotal" dc:"满意度评价总数"`
}

// ComplaintTypesReq 获取投诉类型分布
type ComplaintTypesReq struct {
	g.Meta `path:"/complaint/types" tags:"投诉管理" method:"get" summary:"获取投诉类型分布"`
}

// ComplaintTypesRes 投诉类型分布响应
type ComplaintTypesRes struct {
	Data []*ComplaintTypeData `json:"data" dc:"类型分布数据"`
}

// ComplaintTypeData 投诉类型数据
type ComplaintTypeData struct {
	Type       string `json:"type" dc:"投诉类型"`
	Count      int    `json:"count" dc:"数量"`
	Percentage int    `json:"percentage" dc:"百分比"`
	Trend      string `json:"trend" dc:"趋势(up/down/stable)"`
}

// ComplaintMonthlyTrendsReq 获取月度趋势
type ComplaintMonthlyTrendsReq struct {
	g.Meta `path:"/complaint/monthly-trends" tags:"投诉管理" method:"get" summary:"获取月度趋势"`
}

// ComplaintMonthlyTrendsRes 月度趋势响应
type ComplaintMonthlyTrendsRes struct {
	Data []*ComplaintMonthlyData `json:"data" dc:"月度趋势数据"`
}

// ComplaintMonthlyData 月度数据
type ComplaintMonthlyData struct {
	Month          string `json:"month" dc:"月份"`
	CompletionRate int    `json:"completionRate" dc:"完成率"`
	TotalCount     int    `json:"totalCount" dc:"总数量"`
	CompletedCount int    `json:"completedCount" dc:"已完成数量"`
}

// ComplaintAreasReq 获取区域分布
type ComplaintAreasReq struct {
	g.Meta `path:"/complaint/areas" tags:"投诉管理" method:"get" summary:"获取区域分布"`
}

// ComplaintAreasRes 区域分布响应
type ComplaintAreasRes struct {
	Data []*ComplaintAreaData `json:"data" dc:"区域分布数据"`
}

// ComplaintAreaData 区域数据
type ComplaintAreaData struct {
	Area       string `json:"area" dc:"区域"`
	Count      int    `json:"count" dc:"数量"`
	Percentage int    `json:"percentage" dc:"百分比"`
}
