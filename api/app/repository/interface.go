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

	GetCalendarEvents(ctx context.Context, msg *GetCalendarEvents) ([]CalendarEvent, *int64, error)

	GetSickLeave(ctx context.Context, msg *GetSickLeave) (*SickLeave, error)
	CreateSickLeave(ctx context.Context, msg *CreateSickLeave) (*CreatedSickLeave, error)
	DeleteSickLeave(ctx context.Context, msg *DeleteSickLeave) error

	GetVacation(ctx context.Context, msg *GetVacation) (*Vacation, error)
	CreateVacation(ctx context.Context, msg *CreateVacation) (*CreatedVacation, error)
	DeleteVacation(ctx context.Context, msg *DeleteVacation) error

	GetReport(ctx context.Context, msg *GetReport) (*Report, error)
	GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error)
	CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error)
	UpdateReport(ctx context.Context, msg *UpdateReport) error
	DeleteReport(ctx context.Context, msg *DeleteReport) error
	GetSystemUser(ctx context.Context) (*SystemUser, error)

	GetStatistics(ctx context.Context, msg *GetStatistics) (*Statistics, error)

	AddObjectToUserPermission(ctx context.Context, msg *AddObjectToUserPermission) error
	RemoveObjectFromUserPermission(ctx context.Context, msg *RemoveObjectFromUserPermission) error
}
