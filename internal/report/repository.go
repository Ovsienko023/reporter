package report

import (
	"context"
)

type Repository interface {
	GetReport(ctx context.Context, msg *GetReport) (*Report, error)
	GetReports(ctx context.Context, msg *GetReports) (*Reports, error)
	CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error)
	UpdateReport(ctx context.Context, msg *UpdateReport) error
	DeleteReport(ctx context.Context, msg *DeleteReport) error

	GetSystemUser() *SystemUser
}
