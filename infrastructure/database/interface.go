package database

import (
	"context"
)

type InterfaceDatabase interface {
	GetReport(ctx context.Context, msg *GetReport) (*Report, error)
	GetReports(ctx context.Context, msg *GetReports) (*Reports, *int, error)
	CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error)
	UpdateReport(ctx context.Context, msg *UpdateReport) error
	DeleteReport(ctx context.Context, msg *DeleteReport) error
	GetSystemUser() *SystemUser
}
