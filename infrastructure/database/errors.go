package database

import "errors"

var (
	ErrReportIdNotFound = errors.New("report id not found")
	ErrReportUpdated    = errors.New("error while updating report")
	ErrCredentials      = errors.New("error credentials")
)

func NewInternalError(err error) error {
	if err == nil {
		return nil
	}
	ErrInternal := errors.New(err.Error())
	return ErrInternal
}
