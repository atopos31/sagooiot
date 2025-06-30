package system

import (
	"context"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// SysComplaint 投诉
var SysComplaint = cSysComplaint{}

type cSysComplaint struct{}

// AddComplaint 投诉添加
func (c *cSysComplaint) AddComplaint(ctx context.Context, req *system.AddComplaintReq) (res *system.AddComplaintRes, err error) {
	var input *model.AddComplaintInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysComplaint().Add(ctx, input)
	return
}

// GetComplaintList 投诉列表
func (c *cSysComplaint) GetComplaintList(ctx context.Context, req *system.ComplaintListReq) (res *system.ComplaintListRes, err error) {
	g.Log().Info(ctx, "ComplaintListReq", req)
	var input *model.ComplaintListDoInput
	if err = gconv.Scan(req, &input); err != nil {
		g.Log().Error(ctx, "转换投诉列表输入参数失败", err)
		return
	}
	out, err := service.SysComplaint().ComplaintList(ctx, input)
	if err != nil {
		g.Log().Error(ctx, "获取投诉列表失败", err)
		return
	}
	res = new(system.ComplaintListRes)
	res.Total = out.Total
	res.CurrentPage = out.CurrentPage
	res.Data = out.Data
	return
}

// GetComplaintByTicketNo 根据TicketNo获取投诉信息
func (c *cSysComplaint) GetComplaintByTicketNo(ctx context.Context, req *system.GetComplaintByIdReq) (res *system.GetComplaintByIdRes, err error) {
	out, err := service.SysComplaint().GetComplaintByTicketNo(ctx, req.Id)
	if err != nil {
		return
	}
	var complaintInfoRes *model.ComplaintInfoRes
	if out != nil {
		if err = gconv.Scan(out, &complaintInfoRes); err != nil {
			return
		}
	}
	res = &system.GetComplaintByIdRes{
		Data: complaintInfoRes,
	}
	return
}

// EditComplaint 投诉编辑
func (c *cSysComplaint) EditComplaint(ctx context.Context, req *system.EditComplaintReq) (res *system.EditComplaintRes, err error) {
	var input *model.EditComplaintInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysComplaint().Edit(ctx, input)
	return
}

// DelComplaintByTicketNo 批量删除投诉
func (c *cSysComplaint) DelComplaintByTicketNo(ctx context.Context, req *system.DeleteComplaintByIdReq) (res *system.DeleteComplaintByIdRes, err error) {
	err = service.SysComplaint().DelInfoByTicketNos(ctx, req.Ids)
	return
}

// GetComplaintOverview 获取投诉概要统计
func (c *cSysComplaint) GetComplaintOverview(ctx context.Context, req *system.ComplaintOverviewReq) (res *system.ComplaintOverviewRes, err error) {
	data, err := service.SysComplaint().GetOverview(ctx, req.TimeRange)
	if err != nil {
		return
	}
	res = &system.ComplaintOverviewRes{
		Data: data,
	}
	return
}

// GetComplaintTypes 获取投诉类型分布
func (c *cSysComplaint) GetComplaintTypes(ctx context.Context, req *system.ComplaintTypesReq) (res *system.ComplaintTypesRes, err error) {
	data, err := service.SysComplaint().GetTypes(ctx)
	if err != nil {
		return
	}
	res = &system.ComplaintTypesRes{
		Data: data,
	}
	return
}

// GetComplaintMonthlyTrends 获取月度趋势
func (c *cSysComplaint) GetComplaintMonthlyTrends(ctx context.Context, req *system.ComplaintMonthlyTrendsReq) (res *system.ComplaintMonthlyTrendsRes, err error) {
	data, err := service.SysComplaint().GetMonthlyTrends(ctx)
	if err != nil {
		return
	}
	res = &system.ComplaintMonthlyTrendsRes{
		Data: data,
	}
	return
}

// GetComplaintAreas 获取区域分布
func (c *cSysComplaint) GetComplaintAreas(ctx context.Context, req *system.ComplaintAreasReq) (res *system.ComplaintAreasRes, err error) {
	data, err := service.SysComplaint().GetAreas(ctx)
	if err != nil {
		return
	}
	res = &system.ComplaintAreasRes{
		Data: data,
	}
	return
}

// GetComplaintProcessRecords 获取投诉处理记录
func (c *cSysComplaint) GetComplaintProcessRecords(ctx context.Context, req *system.GetComplaintProcessRecordsReq) (res *system.GetComplaintProcessRecordsRes, err error) {
	// 返回假数据
	mockData := []*system.ComplaintProcessRecord{
		{
			Id:          1,
			TicketNo:    req.TicketNo,
			Status:      "pending",
			Operator:    "系统",
			Action:      "CREATE",
			Description: "投诉工单已创建，等待处理",
			CreatedAt:   "2024/01/15 22:30",
			UpdatedAt:   "2024/01/15 22:30",
		},
		{
			Id:          2,
			TicketNo:    req.TicketNo,
			Status:      "pending",
			Operator:    "王主管",
			Action:      "ASSIGN",
			Description: "已指派给李工程师处理",
			CreatedAt:   "2024/01/15 22:45",
			UpdatedAt:   "2024/01/15 22:45",
		},
		{
			Id:          3,
			TicketNo:    req.TicketNo,
			Status:      "pending",
			Operator:    "李工程师",
			Action:      "START_PROCESS",
			Description: "工程师已接单，正在安排现场勘查",
			CreatedAt:   "2024/01/15 23:20",
			UpdatedAt:   "2024/01/15 23:20",
		},
		{
			Id:          4,
			TicketNo:    req.TicketNo,
			Status:      "pending",
			Operator:    "李工程师",
			Action:      "SITE_SURVEY",
			Description: "已完成现场勘查，确认污水井盖破损，需更换井盖并清理现场",
			CreatedAt:   "2024/01/16 00:30",
			UpdatedAt:   "2024/01/16 00:30",
		},
	}

	res = &system.GetComplaintProcessRecordsRes{
		Data: mockData,
	}
	return
}
