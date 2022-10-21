package database

import "errors"

var (
	ErrInternal         = errors.New("internal error")
	ErrReportIdNotFound = errors.New("report id not found")
	ErrCredentials      = errors.New("error credentials")
)

func NewInternalError(err error) error {
	if err == nil {
		return nil
	}
	ErrInternal := errors.New(err.Error())
	return ErrInternal
}
