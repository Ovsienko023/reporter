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

	message := report.GetReport{
		InvokerId: "11111111-1111-1111-1111-111111111111", // todo
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
	message := report.GetReports{
		InvokerId: "11111111-1111-1111-1111-111111111111",
	}

	result, err := c.repo.GetReports(ctx, &message)
	if err != nil {
		return nil, err
	}

	resp := report.GetReportsResponse{
		Reports: []report.Report{}, //Reports: make([]report.Report, 0),
	}
	for _, obj := range result.Reports {
		resp.Reports = append(resp.Reports, obj)
	}

	return &resp, nil
}
func (c *Core) CreateReport(ctx context.Context, msg *report.CreateReportRequest) (*report.CreatedReportResponse, error) {
	message := report.CreateReport{
		InvokerId: "11111111-1111-1111-1111-111111111111",
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
	message := report.UpdateReport{
		InvokerId: "11111111-1111-1111-1111-111111111111", // todo
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
