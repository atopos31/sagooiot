package model

// AddComplaintInput 添加投诉
type AddComplaintInput struct {
	Title           string `json:"title"`
	Category        string `json:"category"`
	Source          string `json:"source"`
	Area            string `json:"area"`
	ComplainantName string `json:"complainantName"`
	Contact         string `json:"contact"`
	Level           string `json:"level"`
	Content         string `json:"content"`
	Assignee        int    `json:"assignee"`
}

// ComplaintListDoInput 投诉列表
type ComplaintListDoInput struct {
	Name     string
	Status   string
	Category string
	Level    string
	PaginationInput
}

// ComplaintInfoRes 投诉信息响应
type ComplaintInfoRes struct {
	TicketNo        int64   `json:"id" dc:"投诉单号"`
	Title           string  `json:"title" dc:"投诉标题"`
	Category        string  `json:"category" dc:"投诉类别"`
	Source          string  `json:"source" dc:"投诉来源"`
	Area            string  `json:"area" dc:"投诉区域"`
	ComplainantName string  `json:"complainantName" dc:"投诉人姓名"`
	Contact         string  `json:"contact" dc:"联系方式"`
	Level           string  `json:"level" dc:"投诉等级"`
	Content         string  `json:"content" dc:"投诉内容"`
	Assignee        *string `json:"assignee" dc:"指派负责人"`
	Status          string  `json:"status" dc:"投诉状态"`
	CreatedAt       string  `json:"createdAt" dc:"创建时间"`
	UpdatedAt       string  `json:"updatedAt" dc:"更新时间"`
}

// EditComplaintInput 投诉编辑
type EditComplaintInput struct {
	TicketNo        int64  `json:"id" v:"required#投诉单号不能为空" dc:"投诉单号"`
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
