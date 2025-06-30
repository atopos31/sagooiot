package system

import (
	"context"
	"sagooiot/api/v1/common"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// SysComplaintFeedback 投诉反馈
var SysComplaintFeedback = cSysComplaintFeedback{}

type cSysComplaintFeedback struct{}

// AddComplaintFeedback 投诉反馈添加
func (c *cSysComplaintFeedback) AddComplaintFeedback(ctx context.Context, req *system.AddComplaintFeedbackReq) (res *system.AddComplaintFeedbackRes, err error) {
	var input *model.AddComplaintFeedbackInput
	if err = gconv.Scan(req, &input); err != nil {
		g.Log().Error(ctx, "转换投诉反馈添加参数失败:", err)
		return
	}
	err = service.SysComplaintFeedback().Add(ctx, input)
	if err != nil {
		g.Log().Error(ctx, "添加投诉反馈失败:", err)
		return
	}
	res = &system.AddComplaintFeedbackRes{}
	return
}

// GetComplaintFeedbackList 投诉反馈列表
func (c *cSysComplaintFeedback) GetComplaintFeedbackList(ctx context.Context, req *system.ComplaintFeedbackListReq) (res *system.ComplaintFeedbackListRes, err error) {
	g.Log().Info(ctx, "ComplaintFeedbackListReq", req)
	var input *model.ComplaintFeedbackListDoInput
	if err = gconv.Scan(req, &input); err != nil {
		g.Log().Error(ctx, "转换投诉反馈列表输入参数失败:", err)
		return
	}
	out, err := service.SysComplaintFeedback().ComplaintFeedbackList(ctx, input)
	if err != nil {
		g.Log().Error(ctx, "获取投诉反馈列表失败:", err)
		return
	}
	res = &system.ComplaintFeedbackListRes{
		PaginationRes: common.PaginationRes{
			Total:       out.Total,
			CurrentPage: out.CurrentPage,
		},
		Data: out.Data,
	}
	return
}

// DelComplaintFeedbackByIds 批量删除投诉反馈
func (c *cSysComplaintFeedback) DelComplaintFeedbackByIds(ctx context.Context, req *system.DeleteComplaintFeedbackByIdsReq) (res *system.DeleteComplaintFeedbackByIdsRes, err error) {
	err = service.SysComplaintFeedback().DelInfoByIds(ctx, req.Ids)
	if err != nil {
		g.Log().Error(ctx, "批量删除投诉反馈失败:", err)
		return
	}
	res = &system.DeleteComplaintFeedbackByIdsRes{}
	return
}