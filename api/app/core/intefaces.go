package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

type InterfaceCore interface {
	SignUp(ctx context.Context, msg *domain.SignUpRequest) error
	GetToken(ctx context.Context, msg *domain.SignInRequest) (*domain.SignInResponse, error)

	GetProfile(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetProfileResponse, error)

	GetReport(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetReportResponse, error)
	GetReports(ctx context.Context, msg *domain.GetReportsRequest) (*domain.GetReportsResponse, error)
	CreateReport(ctx context.Context, msg *domain.CreateReportRequest) (*domain.CreateReportResponse, error)
	UpdateReport(ctx context.Context, msg *domain.UpdateReportRequest) error
	DeleteReport(ctx context.Context, msg *domain.DeleteReportRequest) error
}
