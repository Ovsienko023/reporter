package core

import "errors"

var (
	ErrInternal            = errors.New("internal error")
	ErrUnauthorized        = errors.New("unauthorized error")
	ErrReportIdNotFound    = errors.New("report id not found")
	ErrSickLeaveIdNotFound = errors.New("sick leave id not found")
	ErrPermissionDenied    = errors.New("permission denied")
	ErrCredentials         = errors.New("permission denied")
	ErrLoginAlreadyInUse   = errors.New("login already in use")
)
