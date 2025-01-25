package errors

import "github.com/pkg/errors"

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrQuotaExceeded   = errors.New("quota exceeded")
	ErrInternalServer  = errors.New("internal server error")
	ErrDuplicateEntry  = errors.New("duplicate entry")
)
