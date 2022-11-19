package repository

import "errors"

var (
	ErrInternal          = errors.New("internal error")
	ErrReportIdNotFound  = errors.New("report id not found")
	ErrCredentials       = errors.New("error credentials")
	ErrLoginAlreadyInUse = errors.New("login already in use")
)

func NewInternalError(err error) error {
	if err == nil {
		return nil
	}
	ErrInternal := errors.New(err.Error())
	return ErrInternal
}
