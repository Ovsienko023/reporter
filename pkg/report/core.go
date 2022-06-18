package report

import "context"

type Core interface {
	GetReport(ctx context.Context, msg *GetReportRequest) (*Report, error)
	GetReports(ctx context.Context, msg *GetReportsRequest) (*Reports, error)
	CreateReport(ctx context.Context, msg *CreateReportRequest) (*CreatedReport, error)
}
