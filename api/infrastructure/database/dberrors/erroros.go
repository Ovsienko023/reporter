package dberrors

import "errors"

var (
	ErrInternal       = errors.New("internal error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("permission denied")
	ErrUserIdNotFound = errors.New("user id not found")
)
