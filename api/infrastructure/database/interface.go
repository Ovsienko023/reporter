package database

import (
	"context"
)

type InterfaceDatabase interface {
	SignUp(ctx context.Context, msg *SignUp) error
	GetAuthUser(ctx context.Context, msg *GetAuthUser) (*Auth, error)

	GetReport(ctx context.Context, msg *GetReport) (*Report, error)
	GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error)
	CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error)
	UpdateReport(ctx context.Context, msg *UpdateReport) error
	DeleteReport(ctx context.Context, msg *DeleteReport) error
	GetSystemUser() *SystemUser
}
