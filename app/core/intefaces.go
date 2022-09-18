package core

import (
	"context"
	domain2 "github.com/Ovsienko023/reporter/app/domain"
)

type InterfaceCore interface {
	GetReport(ctx context.Context, msg *domain2.GetReportRequest) (*domain2.GetReportResponse, error)
	GetReports(ctx context.Context, msg *domain2.GetReportsRequest) (*domain2.GetReportsResponse, error)
	CreateReport(ctx context.Context, msg *domain2.CreateReportRequest) (*domain2.CreateReportResponse, error)
	UpdateReport(ctx context.Context, msg *domain2.UpdateReportRequest) error
	DeleteReport(ctx context.Context, msg *domain2.DeleteReportRequest) error
}
