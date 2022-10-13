package core

import "errors"

var (
	ErrInternal         = errors.New("internal error")
	ErrReportIdNotFound = errors.New("report id not found")
)
