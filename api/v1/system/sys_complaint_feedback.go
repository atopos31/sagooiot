package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// AddComplaintFeedbackReq 添加投诉反馈请求
type AddComplaintFeedbackReq struct {
	g.Meta           `path:"/complaintFeedback" method:"post" summary:"添加投诉反馈" tags:"投诉反馈管理"`
	SurveyCode       string `json:"surveyCode" v:"required#问卷编号不能为空" dc:"问卷编号"`
	TicketNo         int64  `json:"ticketNo" v:"required#投诉标识不能为空" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" v:"required#调查者姓名不能为空" dc:"调查者姓名"`
	ContactInfo      string `json:"contactInfo" v:"required#联系方式不能为空" dc:"联系方式"`
	ProcessingSpeed  string    `json:"processingSpeed" v:"required#处理速度不能为空" dc:"处理速度"`
	StaffAttitude    string    `json:"staffAttitude" v:"required#处理人员态度不能为空" dc:"处理人员态度"`
	ResolutionEffect string    `json:"resolutionEffect" v:"required#问题解决效果不能为空" dc:"问题解决效果"`
	OtherSuggestions string `json:"otherSuggestions" dc:"其它建议或意见"`
}

type AddComplaintFeedbackRes struct{}

// ComplaintFeedbackListReq 投诉反馈列表请求
type ComplaintFeedbackListReq struct {
	g.Meta           `path:"/complaintFeedback/list" method:"get" summary:"投诉反馈列表" tags:"投诉反馈管理"`
	TicketNo         int64 `json:"ticketNo" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" dc:"调查者姓名"`
	SurveyCode       string `json:"surveyCode" dc:"问卷编号"`
	common.PaginationReq
}

type ComplaintFeedbackListRes struct {
	common.PaginationRes
	Data []*model.ComplaintFeedbackInfoRes `json:"Data" dc:"数据列表"`
}

// GetComplaintFeedbackByIdReq 根据ID获取投诉反馈请求
type GetComplaintFeedbackByIdReq struct {
	g.Meta `path:"/complaintFeedback/{id}" method:"get" summary:"根据ID获取投诉反馈" tags:"投诉反馈管理"`
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"投诉反馈ID"`
}

type GetComplaintFeedbackByIdRes struct {
	Data *model.ComplaintFeedbackInfoRes `json:"data" dc:"投诉反馈信息"`
}

// EditComplaintFeedbackReq 编辑投诉反馈请求
type EditComplaintFeedbackReq struct {
	g.Meta           `path:"/complaintFeedback/{id}" method:"put" summary:"编辑投诉反馈" tags:"投诉反馈管理"`
	Id               int64  `json:"id" v:"required#ID不能为空" dc:"投诉反馈ID"`
	SurveyCode       string `json:"surveyCode" v:"required#问卷编号不能为空" dc:"问卷编号"`
	TicketNo         int64 `json:"ticketNo" v:"required#投诉标识不能为空" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" v:"required#调查者姓名不能为空" dc:"调查者姓名"`
	ContactInfo      string `json:"contactInfo" v:"required#联系方式不能为空" dc:"联系方式"`
	ProcessingSpeed  int    `json:"processingSpeed" v:"required|between:1,5#处理速度不能为空|处理速度必须在1-5之间" dc:"处理速度"`
	StaffAttitude    int    `json:"staffAttitude" v:"required|between:1,5#处理人员态度不能为空|处理人员态度必须在1-5之间" dc:"处理人员态度"`
	ResolutionEffect int    `json:"resolutionEffect" v:"required|between:1,5#问题解决效果不能为空|问题解决效果必须在1-5之间" dc:"问题解决效果"`
	OtherSuggestions string `json:"otherSuggestions" dc:"其它建议或意见"`
}

type EditComplaintFeedbackRes struct{}

// DeleteComplaintFeedbackByIdReq 根据ID删除投诉反馈请求
type DeleteComplaintFeedbackByIdReq struct {
	g.Meta `path:"/complaintFeedback/{id}" method:"delete" summary:"根据ID删除投诉反馈" tags:"投诉反馈管理"`
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"投诉反馈ID"`
}

type DeleteComplaintFeedbackByIdRes struct{}

// DeleteComplaintFeedbackByIdsReq 批量删除投诉反馈请求
type DeleteComplaintFeedbackByIdsReq struct {
	g.Meta `path:"/complaintFeedback/batch" method:"delete" summary:"批量删除投诉反馈" tags:"投诉反馈管理"`
	Ids    []int64 `json:"ids" v:"required#ID列表不能为空" dc:"投诉反馈ID列表"`
}

type DeleteComplaintFeedbackByIdsRes struct{}
