package system

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"sagooiot/api/v1/system"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
)

type sSysComplaint struct {
}

func sysComplaintNew() *sSysComplaint {
	return &sSysComplaint{}
}

func init() {
	service.RegisterSysComplaint(sysComplaintNew())
}

// getRoleName 根据角色ID获取角色名称
func (s *sSysComplaint) getRoleName(ctx context.Context, roleId int) *string {
	if roleId == 0 {
		return nil
	}

	var role *entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, roleId).Scan(&role)
	if err != nil || role == nil {
		return nil
	}

	return &role.UserNickname
}

// Add 添加投诉
func (s *sSysComplaint) Add(ctx context.Context, in *model.AddComplaintInput) (err error) {
	// 生成8位随机ticket_no
	ticketNo := gtime.Now().Layout("0601") + gconv.String(grand.Intn(900000)+100000) // 生成10000000-99999999之间的随机数
	// 构建投诉数据
	complaintData := &do.SysComplaints{
		TicketNo:           ticketNo,
		Title:              in.Title,
		Content:            in.Content,
		TypeCode:           in.Category,
		SourceCode:         in.Source,
		PriorityCode:       in.Level,
		StatusCode:         "pending", // 默认状态为待处理
		ComplainantName:    in.ComplainantName,
		ComplainantContact: in.Contact,
		AreaCode:           in.Area,
		AssigneeId:         in.Assignee,
		CreatedAt:          gtime.Now(),
		UpdatedAt:          gtime.Now(),
	}

	// 插入投诉记录
	_, err = dao.SysComplaints.Ctx(ctx).Data(complaintData).Insert()
	if err != nil {
		return err
	}

	// // 记录操作日志
	// logData := &do.SysComplaintLogs{
	// 	ComplaintId:   complaintId,
	// 	ActionType:    "create",
	// 	ActionUser:    "system", // 这里应该从上下文获取当前用户
	// 	ActionContent: "创建投诉",
	// 	CreatedAt:     gtime.Now(),
	// }
	// _, err = dao.SysComplaintLogs.Ctx(ctx).Data(logData).Insert()

	return err
}

// ComplaintList 投诉列表
func (s *sSysComplaint) ComplaintList(ctx context.Context, input *model.ComplaintListDoInput) (out *model.ComplaintListOutput, err error) {
	m := dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().IsDeleted, 0)

	// 构建查询条件
	if input.Name != "" {
		m = m.WhereLike(dao.SysComplaints.Columns().Title, "%"+input.Name+"%")
	}
	if input.Status != "" {
		m = m.Where(dao.SysComplaints.Columns().StatusCode, input.Status)
	}
	if input.Category != "" {
		m = m.Where(dao.SysComplaints.Columns().TypeCode, input.Category)
	}
	if input.Level != "" {
		m = m.Where(dao.SysComplaints.Columns().PriorityCode, input.Level)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return
	}

	// 分页查询
	var entities []*entity.SysComplaints
	err = m.Page(input.PageNum, input.PageSize).OrderDesc(dao.SysComplaints.Columns().CreatedAt).Scan(&entities)
	if err != nil {
		return
	}

	// 转换为响应格式
	list := make([]*model.ComplaintInfoRes, len(entities))
	for i, entity := range entities {
		// 获取角色名称
		assigneeName := s.getRoleName(ctx, entity.AssigneeId)

		list[i] = &model.ComplaintInfoRes{
			TicketNo:        entity.TicketNo,
			Title:           entity.Title,
			Category:        entity.TypeCode,
			Source:          entity.SourceCode,
			Area:            entity.AreaCode,
			ComplainantName: entity.ComplainantName,
			Contact:         entity.ComplainantContact,
			Level:           entity.PriorityCode,
			Content:         entity.Content,
			Assignee:        assigneeName,
			Status:          entity.StatusCode,
			CreatedAt:       entity.CreatedAt.String(),
			UpdatedAt:       entity.UpdatedAt.String(),
		}
	}

	out = &model.ComplaintListOutput{
		PaginationOutput: model.PaginationOutput{
			Total:       total,
			CurrentPage: input.PageNum,
		},
		Data: list,
	}
	return
}

// GetComplaintByTicketNo 根据TicketNo获取投诉信息
func (s *sSysComplaint) GetComplaintByTicketNo(ctx context.Context, ticketNo int64) (out *model.ComplaintInfoRes, err error) {
	var entity *entity.SysComplaints
	err = dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().TicketNo, ticketNo).Where(dao.SysComplaints.Columns().IsDeleted, 0).Scan(&entity)
	if err != nil {
		return
	}
	if entity == nil {
		return nil, fmt.Errorf("投诉信息不存在")
	}

	// 获取角色名称
	assigneeName := s.getRoleName(ctx, entity.AssigneeId)

	out = &model.ComplaintInfoRes{
		TicketNo:        entity.TicketNo,
		Title:           entity.Title,
		Category:        entity.TypeCode,
		Source:          entity.SourceCode,
		Area:            entity.AreaCode,
		ComplainantName: entity.ComplainantName,
		Contact:         entity.ComplainantContact,
		Level:           entity.PriorityCode,
		Content:         entity.Content,
		Assignee:        assigneeName,
		Status:          entity.StatusCode,
		CreatedAt:       entity.CreatedAt.String(),
		UpdatedAt:       entity.UpdatedAt.String(),
	}

	return
}

// Edit 编辑投诉
func (s *sSysComplaint) Edit(ctx context.Context, in *model.EditComplaintInput) (err error) {
	// 构建更新数据
	updateData := &do.SysComplaints{
		Title:              in.Title,
		Content:            in.Content,
		TypeCode:           in.Category,
		SourceCode:         in.Source,
		PriorityCode:       in.Level,
		ComplainantName:    in.ComplainantName,
		ComplainantContact: in.Contact,
		AreaCode:           in.Area,
		AssigneeId:         in.Assignee,
		UpdatedAt:          gtime.Now(),
	}

	// 更新投诉记录
	_, err = dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().TicketNo, in.TicketNo).Data(updateData).Update()
	if err != nil {
		return err
	}

	// // 记录操作日志
	// logData := &do.SysComplaintLogs{
	// 	ComplaintId:   gconv.String(in.TicketNo),
	// 	ActionType:    "update",
	// 	ActionUser:    "system", // 这里应该从上下文获取当前用户
	// 	ActionContent: "编辑投诉信息",
	// 	CreatedAt:     gtime.Now(),
	// }
	// _, err = dao.SysComplaintLogs.Ctx(ctx).Data(logData).Insert()

	return err
}

// DelInfoByTicketNo 根据TicketNo删除投诉（软删除）
func (s *sSysComplaint) DelInfoByTicketNo(ctx context.Context, ticketNo int64) (err error) {
	// 软删除投诉记录
	updateData := &do.SysComplaints{
		IsDeleted: 1,
		DeletedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}

	_, err = dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().TicketNo, ticketNo).Data(updateData).Update()
	if err != nil {
		return err
	}

	// 记录操作日志
	logData := &do.SysComplaintLogs{
		ComplaintId:   gconv.String(ticketNo),
		ActionType:    "delete",
		ActionUser:    "system", // 这里应该从上下文获取当前用户
		ActionContent: "删除投诉",
		CreatedAt:     gtime.Now(),
	}
	_, err = dao.SysComplaintLogs.Ctx(ctx).Data(logData).Insert()

	return err
}

// DelInfoByTicketNos 批量删除投诉（软删除）
func (s *sSysComplaint) DelInfoByTicketNos(ctx context.Context, ticketNos []int64) (err error) {
	if len(ticketNos) == 0 {
		return fmt.Errorf("TicketNo列表不能为空")
	}

	// 批量软删除投诉记录
	updateData := &do.SysComplaints{
		IsDeleted: 1,
		DeletedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}

	_, err = dao.SysComplaints.Ctx(ctx).WhereIn(dao.SysComplaints.Columns().TicketNo, ticketNos).Data(updateData).Update()
	if err != nil {
		return err
	}

	// 批量记录操作日志
	// for _, ticketNo := range ticketNos {
	// 	logData := &do.SysComplaintLogs{
	// 		ComplaintId:   gconv.String(ticketNo),
	// 		ActionType:    "delete",
	// 		ActionUser:    "system", // 这里应该从上下文获取当前用户
	// 		ActionContent: "批量删除投诉",
	// 		CreatedAt:     gtime.Now(),
	// 	}
	// 	_, err = dao.SysComplaintLogs.Ctx(ctx).Data(logData).Insert()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// GetOverview 获取投诉概要统计
func (s *sSysComplaint) GetOverview(ctx context.Context, timeRange string) (out *system.ComplaintOverviewData, err error) {
	// 构建时间范围查询条件
	m := dao.SysComplaints.Ctx(ctx).Where(dao.SysComplaints.Columns().IsDeleted, 0)
	if timeRange != "" {
		switch timeRange {
		case "today":
			m = m.WhereGTE(dao.SysComplaints.Columns().CreatedAt, gtime.Now().StartOfDay())
		case "week":
			m = m.WhereGTE(dao.SysComplaints.Columns().CreatedAt, gtime.Now().AddDate(0, 0, -7))
		case "month":
			m = m.WhereGTE(dao.SysComplaints.Columns().CreatedAt, gtime.Now().AddDate(0, -1, 0))
		case "year":
			m = m.WhereGTE(dao.SysComplaints.Columns().CreatedAt, gtime.Now().AddDate(-1, 0, 0))
		}
	}

	// 获取总投诉数
	totalComplaints, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 获取各状态投诉数
	pendingCount, _ := m.Where(dao.SysComplaints.Columns().StatusCode, "pending").Count()
	completedCount, _ := m.Where(dao.SysComplaints.Columns().StatusCode, "completed").Count()
	urgentCount, _ := m.Where(dao.SysComplaints.Columns().PriorityCode, "urgent").Count()

	// 计算完成率
	var completionRate float64
	if totalComplaints > 0 {
		completionRate = float64(completedCount) / float64(totalComplaints) * 100
	}

	out = &system.ComplaintOverviewData{
		TotalComplaints:       totalComplaints,
		PendingComplaints:     pendingCount,
		CompletedComplaints:   completedCount,
		UrgentComplaints:      urgentCount,
		AverageProcessingTime: 2.3,
		CompletionRate:        completionRate,
		SatisfactionScore:     4.2,
		SatisfactionTotal:     totalComplaints,
	}
	return
}

// GetTypes 获取投诉类型分布
func (s *sSysComplaint) GetTypes(ctx context.Context) (out []*system.ComplaintTypeData, err error) {
	// 查询投诉类型分布
	type TypeCount struct {
		TypeCode string `json:"type_code"`
		Count    int    `json:"count"`
	}

	var typeCounts []TypeCount
	err = dao.SysComplaints.Ctx(ctx).
		Where(dao.SysComplaints.Columns().IsDeleted, 0).
		Fields("type_code, COUNT(*) as count").
		Group("type_code").
		OrderDesc("count").
		Scan(&typeCounts)
	if err != nil {
		return nil, err
	}

	// 计算总数用于百分比计算
	totalCount := 0
	for _, tc := range typeCounts {
		totalCount += tc.Count
	}

	// 转换为响应格式
	out = make([]*system.ComplaintTypeData, len(typeCounts))
	for i, tc := range typeCounts {
		percentage := 0
		if totalCount > 0 {
			percentage = tc.Count * 100 / totalCount
		}
		out[i] = &system.ComplaintTypeData{
			Type:       tc.TypeCode,
			Count:      tc.Count,
			Percentage: percentage,
			Trend:      "stable",
		}
	}

	return
}

// GetMonthlyTrends 获取月度趋势
func (s *sSysComplaint) GetMonthlyTrends(ctx context.Context) (out []*system.ComplaintMonthlyData, err error) {
	// 获取12个月前的时间
	twelveMonthsAgo := gtime.Now().AddDate(0, -12, 0)

	// 查询最近12个月的投诉数据
	var complaints []*entity.SysComplaints
	err = dao.SysComplaints.Ctx(ctx).
		Where(dao.SysComplaints.Columns().IsDeleted, 0).
		WhereGTE(dao.SysComplaints.Columns().CreatedAt, twelveMonthsAgo).
		OrderAsc(dao.SysComplaints.Columns().CreatedAt).
		Scan(&complaints)
	if err != nil {
		return nil, err
	}

	// 按月份统计数据
	monthlyStats := make(map[string]*system.ComplaintMonthlyData)

	for _, complaint := range complaints {
		if complaint.CreatedAt == nil {
			continue
		}

		// 格式化月份
		month := complaint.CreatedAt.Layout("2006-01")

		// 初始化月份数据
		if monthlyStats[month] == nil {
			monthlyStats[month] = &system.ComplaintMonthlyData{
				Month:          month,
				TotalCount:     0,
				CompletedCount: 0,
				CompletionRate: 0,
			}
		}

		// 统计总数
		monthlyStats[month].TotalCount++

		// 统计已完成数量
		if complaint.StatusCode == "completed" {
			monthlyStats[month].CompletedCount++
		}
	}

	// 计算完成率并转换为切片
	out = make([]*system.ComplaintMonthlyData, 0, len(monthlyStats))
	for _, data := range monthlyStats {
		if data.TotalCount > 0 {
			data.CompletionRate = data.CompletedCount * 100 / data.TotalCount
		}
		out = append(out, data)
	}

	// 按月份排序
	for i := 0; i < len(out)-1; i++ {
		for j := i + 1; j < len(out); j++ {
			if out[i].Month > out[j].Month {
				out[i], out[j] = out[j], out[i]
			}
		}
	}

	return
}

// GetAreas 获取区域分布
func (s *sSysComplaint) GetAreas(ctx context.Context) (out []*system.ComplaintAreaData, err error) {
	// 查询所有未删除的投诉数据
	var complaints []*entity.SysComplaints
	err = dao.SysComplaints.Ctx(ctx).
		Where(dao.SysComplaints.Columns().IsDeleted, 0).
		Fields(dao.SysComplaints.Columns().AreaCode).
		Scan(&complaints)
	if err != nil {
		return nil, err
	}

	// 统计各区域的投诉数量
	areaStats := make(map[string]int)
	totalCount := 0

	for _, complaint := range complaints {
		if complaint.AreaCode != "" {
			areaStats[complaint.AreaCode]++
			totalCount++
		}
	}

	// 转换为响应格式并按数量排序
	type areaData struct {
		area  string
		count int
	}

	areaList := make([]areaData, 0, len(areaStats))
	for area, count := range areaStats {
		areaList = append(areaList, areaData{
			area:  area,
			count: count,
		})
	}

	// 按数量降序排序
	for i := 0; i < len(areaList)-1; i++ {
		for j := i + 1; j < len(areaList); j++ {
			if areaList[i].count < areaList[j].count {
				areaList[i], areaList[j] = areaList[j], areaList[i]
			}
		}
	}

	// 转换为最终响应格式
	out = make([]*system.ComplaintAreaData, len(areaList))
	for i, ad := range areaList {
		percentage := 0
		if totalCount > 0 {
			percentage = ad.count * 100 / totalCount
		}
		out[i] = &system.ComplaintAreaData{
			Area:       ad.area,
			Count:      ad.count,
			Percentage: percentage,
		}
	}

	return
}
