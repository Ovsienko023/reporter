package core

import (
	"context"
	"github.com/Ovsienko023/reporter/internal/report"
)

type Core struct {
	repo report.Repository
}

func NewCore(r report.Repository) *Core {
	return &Core{
		repo: r,
	}
}

func (c *Core) GetReport(ctx context.Context, msg *report.GetReportRequest) (*report.GetReportResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := report.GetReport{
		InvokerId: *systemUser.UserId,
		ReportId:  msg.ReportId,
	}

	result, err := c.repo.GetReport(ctx, &message)
	if err != nil {
		return nil, err
	}

	resp := report.GetReportResponse{
		Report: *result,
	}

	return &resp, nil
}

func (c *Core) GetReports(ctx context.Context, msg *report.GetReportsRequest) (*report.GetReportsResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := report.GetReports{
		InvokerId: *systemUser.UserId,
	}

	result, err := c.repo.GetReports(ctx, &message)
	if err != nil {
		return nil, err
	}

	resp := report.GetReportsResponse{
		Reports: []report.Report{},
	}

	for _, obj := range result.Reports {
		resp.Reports = append(resp.Reports, obj)
	}

	return &resp, nil
}
func (c *Core) CreateReport(ctx context.Context, msg *report.CreateReportRequest) (*report.CreatedReportResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := report.CreateReport{
		InvokerId: *systemUser.UserId,
		Title:     msg.Title,
		StartTime: msg.StartTime,
		EndTime:   msg.EndTime,
		BreakTime: msg.BreakTime,
		WorkTime:  msg.WorkTime,
		Body:      msg.Body,
	}
	result, err := c.repo.CreateReport(ctx, &message)
	if err != nil {
		return nil, err
	}

	resp := report.CreatedReportResponse{
		Id: result.Id,
	}

	return &resp, nil
}

func (c *Core) UpdateReport(ctx context.Context, msg *report.UpdateReportRequest) error {
	systemUser := c.repo.GetSystemUser()

	message := report.UpdateReport{
		InvokerId: *systemUser.UserId,
		ReportId:  msg.ReportId,
		Title:     msg.Title,
		StartTime: msg.StartTime,
		EndTime:   msg.EndTime,
		BreakTime: msg.BreakTime,
		WorkTime:  msg.WorkTime,
		Body:      msg.Body,
	}

	err := c.repo.UpdateReport(ctx, &message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) DeleteReport(ctx context.Context, msg *report.DeleteReportRequest) error {
	systemUser := c.repo.GetSystemUser()

	message := report.DeleteReport{
		InvokerId: *systemUser.UserId,
		ReportId:  msg.ReportId,
	}

	err := c.repo.DeleteReport(ctx, &message)
	if err != nil {
		return err
	}

	return nil
}
