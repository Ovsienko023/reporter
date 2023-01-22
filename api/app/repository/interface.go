package repository

import (
	"context"
)

type InterfaceDatabase interface {
	SignUp(ctx context.Context, msg *SignUp) error
	SignIn(ctx context.Context, msg *SignIn) (*Auth, error)

	GetProfile(ctx context.Context, msg *GetProfile) (*Profile, error)
	UpdateProfile(ctx context.Context, msg *UpdateProfile) error

	GetUsers(ctx context.Context, msg *GetUsers) ([]UserItem, *int, error)

	GetReport(ctx context.Context, msg *GetReport) (*Report, error)
	GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error)
	CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error)
	UpdateReport(ctx context.Context, msg *UpdateReport) error
	DeleteReport(ctx context.Context, msg *DeleteReport) error
	GetSystemUser() *SystemUser

	GetStatistics(ctx context.Context, msg *GetStatistics) (*Statistics, error)
}
