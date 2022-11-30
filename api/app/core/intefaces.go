package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

type InterfaceCore interface {
	SignIn(ctx context.Context, msg *domain.SignInRequest) (*domain.SignInResponse, error)
	SignUp(ctx context.Context, msg *domain.SignUpRequest) error

	SendEmail(ctx context.Context, msg *domain.SendEmailRequest) error

	GetProfile(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetProfileResponse, error)

	GetReport(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetReportResponse, error)
	GetReports(ctx context.Context, msg *domain.GetReportsRequest) (*domain.GetReportsResponse, error)
	CreateReport(ctx context.Context, msg *domain.CreateReportRequest) (*domain.CreateReportResponse, error)
	UpdateReport(ctx context.Context, msg *domain.UpdateReportRequest) error
	DeleteReport(ctx context.Context, msg *domain.DeleteReportRequest) error
}
