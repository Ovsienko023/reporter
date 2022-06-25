package report

import "context"

type Core interface {
	GetReport(ctx context.Context, msg *GetReportRequest) (*GetReportResponse, error)
	GetReports(ctx context.Context, msg *GetReportsRequest) (*GetReportsResponse, error)
	CreateReport(ctx context.Context, msg *CreateReportRequest) (*CreatedReportResponse, error)
}
