package system

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
)

type sSysComplaintFeedback struct {
}

func sysComplaintFeedbackNew() *sSysComplaintFeedback {
	return &sSysComplaintFeedback{}
}

func init() {
	service.RegisterSysComplaintFeedback(sysComplaintFeedbackNew())
}

// Add 添加投诉反馈
func (s *sSysComplaintFeedback) Add(ctx context.Context, in *model.AddComplaintFeedbackInput) (err error) {
	// 检查投诉记录是否存在
	var count int
	count, err = dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().TicketNo, in.TicketNo).Where(dao.SysComplaints.Columns().IsDeleted, 0).Count()
	if err != nil {
		g.Log().Error(ctx, "检查投诉记录失败:", err)
		return fmt.Errorf("检查投诉记录失败: %v", err)
	}
	if count == 0 {
		return fmt.Errorf("投诉记录不存在，TicketNo: %s", in.TicketNo)
	}
	// 构建投诉反馈数据
	data := do.SysComplaintFeedback{
		SurveyCode:       in.SurveyCode,
		TicketNo:         in.TicketNo,
		InvestigatorName: in.InvestigatorName,
		ContactInfo:      in.ContactInfo,
		ProcessingSpeed:  in.ProcessingSpeed,
		StaffAttitude:    in.StaffAttitude,
		ResolutionEffect: in.ResolutionEffect,
		OtherSuggestions: in.OtherSuggestions,
		CreatedAt:        gtime.Now(),
		IsDeleted:        0,
	}

	// 插入数据
	_, err = dao.SysComplaintFeedback.Ctx(ctx).Data(data).Insert()
	if err != nil {
		g.Log().Error(ctx, "添加投诉反馈失败:", err)
		return fmt.Errorf("添加投诉反馈失败: %v", err)
	}

	return nil
}

// ComplaintFeedbackList 投诉反馈列表
func (s *sSysComplaintFeedback) ComplaintFeedbackList(ctx context.Context, input *model.ComplaintFeedbackListDoInput) (total int, list []*model.ComplaintFeedbackInfoRes, err error) {
	m := dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().IsDeleted, 0)

	// 条件查询
	if input.TicketNo != "" {
		m = m.WhereLike(dao.SysComplaintFeedback.Columns().TicketNo, "%"+input.TicketNo+"%")
	}
	if input.InvestigatorName != "" {
		m = m.WhereLike(dao.SysComplaintFeedback.Columns().InvestigatorName, "%"+input.InvestigatorName+"%")
	}
	if input.SurveyCode != "" {
		m = m.WhereLike(dao.SysComplaintFeedback.Columns().SurveyCode, "%"+input.SurveyCode+"%")
	}

	// 获取总数
	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, "获取投诉反馈总数失败:", err)
		return 0, nil, err
	}

	// 分页查询
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = 10
	}

	var feedbacks []*entity.SysComplaintFeedback
	err = m.Page(input.PageNum, input.PageSize).OrderDesc(dao.SysComplaintFeedback.Columns().CreatedAt).Scan(&feedbacks)
	if err != nil {
		g.Log().Error(ctx, "获取投诉反馈列表失败:", err)
		return 0, nil, err
	}

	// 转换为响应格式
	list = make([]*model.ComplaintFeedbackInfoRes, 0, len(feedbacks))
	for _, feedback := range feedbacks {
		item := &model.ComplaintFeedbackInfoRes{
			Id:               feedback.Id,
			SurveyCode:       feedback.SurveyCode,
			TicketNo:         feedback.TicketNo,
			InvestigatorName: feedback.InvestigatorName,
			ContactInfo:      feedback.ContactInfo,
			ProcessingSpeed:  feedback.ProcessingSpeed,
			StaffAttitude:    feedback.StaffAttitude,
			ResolutionEffect: feedback.ResolutionEffect,
			OtherSuggestions: feedback.OtherSuggestions,
			CreatedAt:        feedback.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		list = append(list, item)
	}

	return total, list, nil
}

// GetComplaintFeedbackById 根据ID获取投诉反馈信息
func (s *sSysComplaintFeedback) GetComplaintFeedbackById(ctx context.Context, id int64) (out *model.ComplaintFeedbackInfoRes, err error) {
	var feedback *entity.SysComplaintFeedback
	err = dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().Id, id).Where(dao.SysComplaintFeedback.Columns().IsDeleted, 0).Scan(&feedback)
	if err != nil {
		g.Log().Error(ctx, "获取投诉反馈信息失败:", err)
		return nil, err
	}

	if feedback == nil {
		return nil, fmt.Errorf("投诉反馈信息不存在")
	}

	out = &model.ComplaintFeedbackInfoRes{
		Id:               feedback.Id,
		SurveyCode:       feedback.SurveyCode,
		TicketNo:         feedback.TicketNo,
		InvestigatorName: feedback.InvestigatorName,
		ContactInfo:      feedback.ContactInfo,
		ProcessingSpeed:  feedback.ProcessingSpeed,
		StaffAttitude:    feedback.StaffAttitude,
		ResolutionEffect: feedback.ResolutionEffect,
		OtherSuggestions: feedback.OtherSuggestions,
		CreatedAt:        feedback.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return out, nil
}

// Edit 编辑投诉反馈
func (s *sSysComplaintFeedback) Edit(ctx context.Context, in *model.EditComplaintFeedbackInput) (err error) {
	// 检查投诉反馈记录是否存在
	count, err := dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().Id, in.TicketNo).Where(dao.SysComplaintFeedback.Columns().IsDeleted, 0).Count()
	if err != nil {
		g.Log().Error(ctx, "检查投诉反馈记录失败:", err)
		return err
	}
	if count == 0 {
		return fmt.Errorf("投诉反馈记录不存在")
	}

	// 检查投诉记录是否存在
	var complaintCount int
	complaintCount, err = dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().TicketNo, in.TicketNo).Where(dao.SysComplaints.Columns().IsDeleted, 0).Count()
	if err != nil {
		g.Log().Error(ctx, "检查投诉记录失败:", err)
		return fmt.Errorf("检查投诉记录失败: %v", err)
	}
	if complaintCount == 0 {
		return fmt.Errorf("投诉记录不存在，TicketNo: %s", in.TicketNo)
	}

	// 检查是否已存在该TicketNo的其他反馈记录（排除当前记录）
	var feedbackCount int
	feedbackCount, err = dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().TicketNo, in.TicketNo).Where(dao.SysComplaintFeedback.Columns().IsDeleted, 0).WhereNot(dao.SysComplaintFeedback.Columns().Id, in.Id).Count()
	if err != nil {
		g.Log().Error(ctx, "检查投诉反馈记录失败:", err)
		return fmt.Errorf("检查投诉反馈记录失败: %v", err)
	}
	if feedbackCount > 0 {
		return fmt.Errorf("该投诉已存在其他反馈记录，TicketNo: %s", in.TicketNo)
	}

	// 构建更新数据
	data := do.SysComplaintFeedback{
		SurveyCode:       in.SurveyCode,
		TicketNo:         in.TicketNo,
		InvestigatorName: in.InvestigatorName,
		ContactInfo:      in.ContactInfo,
		ProcessingSpeed:  in.ProcessingSpeed,
		StaffAttitude:    in.StaffAttitude,
		ResolutionEffect: in.ResolutionEffect,
		OtherSuggestions: in.OtherSuggestions,
	}

	// 更新数据
	_, err = dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().Id, in.Id).Data(data).Update()
	if err != nil {
		g.Log().Error(ctx, "编辑投诉反馈失败:", err)
		return fmt.Errorf("编辑投诉反馈失败: %v", err)
	}

	return nil
}

// DelInfoById 根据ID删除投诉反馈（软删除）
func (s *sSysComplaintFeedback) DelInfoById(ctx context.Context, id int64) (err error) {
	// 检查记录是否存在
	count, err := dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().Id, id).Where(dao.SysComplaintFeedback.Columns().IsDeleted, 0).Count()
	if err != nil {
		g.Log().Error(ctx, "检查投诉反馈记录失败:", err)
		return err
	}
	if count == 0 {
		return fmt.Errorf("投诉反馈记录不存在")
	}

	// 软删除
	data := do.SysComplaintFeedback{
		IsDeleted: 1,
		DeletedAt: gtime.Now(),
	}

	_, err = dao.SysComplaintFeedback.Ctx(ctx).Where(dao.SysComplaintFeedback.Columns().Id, id).Data(data).Update()
	if err != nil {
		g.Log().Error(ctx, "删除投诉反馈失败:", err)
		return fmt.Errorf("删除投诉反馈失败: %v", err)
	}

	return nil
}

// DelInfoByIds 批量删除投诉反馈（软删除）
func (s *sSysComplaintFeedback) DelInfoByIds(ctx context.Context, ids []int64) (err error) {
	if len(ids) == 0 {
		return fmt.Errorf("删除ID列表不能为空")
	}

	// 软删除
	data := do.SysComplaintFeedback{
		IsDeleted: 1,
		DeletedAt: gtime.Now(),
	}

	_, err = dao.SysComplaintFeedback.Ctx(ctx).WhereIn(dao.SysComplaintFeedback.Columns().Id, ids).Where(dao.SysComplaintFeedback.Columns().IsDeleted, 0).Data(data).Update()
	if err != nil {
		g.Log().Error(ctx, "批量删除投诉反馈失败:", err)
		return fmt.Errorf("批量删除投诉反馈失败: %v", err)
	}

	return nil
}
