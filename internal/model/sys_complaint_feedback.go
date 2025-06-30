package model

// AddComplaintFeedbackInput 添加投诉反馈
type AddComplaintFeedbackInput struct {
	SurveyCode       string `json:"surveyCode" v:"required#问卷编号不能为空" dc:"问卷编号"`
	TicketNo         int64  `json:"ticketNo" v:"required#投诉标识不能为空" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" v:"required#调查者姓名不能为空" dc:"调查者姓名"`
	ContactInfo      string `json:"contactInfo" v:"required#联系方式不能为空" dc:"联系方式"`
	ProcessingSpeed  string `json:"processingSpeed" v:"required#处理速度不能为空" dc:"处理速度"`
	StaffAttitude    string `json:"staffAttitude" v:"required#处理人员态度不能为空" dc:"处理人员态度"`
	ResolutionEffect string `json:"resolutionEffect" v:"required#问题解决效果不能为空" dc:"问题解决效果"`
	OtherSuggestions string `json:"otherSuggestions" dc:"其它建议或意见"`
}

// ComplaintFeedbackListDoInput 投诉反馈列表查询输入
type ComplaintFeedbackListDoInput struct {
	TicketNo         string `json:"ticketNo" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" dc:"调查者姓名"`
	SurveyCode       string `json:"surveyCode" dc:"问卷编号"`
	PaginationInput
}

// ComplaintFeedbackInfoRes 投诉反馈信息响应
type ComplaintFeedbackInfoRes struct {
	Id               int64  `json:"id" dc:"主键ID"`
	SurveyCode       string `json:"surveyCode" dc:"问卷编号"`
	TicketNo         string  `json:"ticketNo" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" dc:"调查者姓名"`
	ContactInfo      string `json:"contactInfo" dc:"联系方式"`
	ProcessingSpeed  string `json:"processingSpeed" dc:"处理速度"`
	StaffAttitude    string `json:"staffAttitude" dc:"处理人员态度"`
	ResolutionEffect string `json:"resolutionEffect" dc:"问题解决效果"`
	OtherSuggestions string `json:"otherSuggestions" dc:"其它建议或意见"`
	CreatedAt        string `json:"createdAt" dc:"创建时间"`
}

// ComplaintFeedbackListOutput 投诉反馈列表输出
type ComplaintFeedbackListOutput struct {
	PaginationOutput
	Data []*ComplaintFeedbackInfoRes `json:"data" dc:"投诉反馈列表数据"`
}

// EditComplaintFeedbackInput 编辑投诉反馈
type EditComplaintFeedbackInput struct {
	Id               int64  `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
	SurveyCode       string `json:"surveyCode" v:"required#问卷编号不能为空" dc:"问卷编号"`
	TicketNo         string `json:"ticketNo" v:"required#投诉标识不能为空" dc:"投诉标识"`
	InvestigatorName string `json:"investigatorName" v:"required#调查者姓名不能为空" dc:"调查者姓名"`
	ContactInfo      string `json:"contactInfo" v:"required#联系方式不能为空" dc:"联系方式"`
	ProcessingSpeed  int    `json:"processingSpeed" v:"required|between:1,5#处理速度不能为空|处理速度必须在1-5之间" dc:"处理速度"`
	StaffAttitude    int    `json:"staffAttitude" v:"required|between:1,5#处理人员态度不能为空|处理人员态度必须在1-5之间" dc:"处理人员态度"`
	ResolutionEffect int    `json:"resolutionEffect" v:"required|between:1,5#问题解决效果不能为空|问题解决效果必须在1-5之间" dc:"问题解决效果"`
	OtherSuggestions string `json:"otherSuggestions" dc:"其它建议或意见"`
}
