package errors

import "github.com/pkg/errors"

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrInvalidArgument = errors.New("invalid argument")
)
