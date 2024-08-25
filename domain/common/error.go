package domain

import (
	"errors"

	"github.com/route-test/constant/logcode"
)

const (
	badRequest          = "bad_request"
	unauthorized        = "unauthorized"
	forbidden           = "forbidden"
	notFound            = "not_found"
	internalServerError = "internal_server_error"
	serviceUnavailable  = "service_unavailable"
)

var (
	ErrBadRequest          = errors.New(badRequest)
	ErrUnauthorized        = errors.New(unauthorized)
	ErrForbidden           = errors.New(forbidden)
	ErrNotFound            = errors.New(notFound)
	ErrInternalServerError = errors.New(internalServerError)
	ErrServiceUnavailable  = errors.New(internalServerError)
)

type DomainError struct {
	Err       error
	Code      logcode.LogCode
	Message   string
	LocaleKey string
}

func (e *DomainError) Error() string {
	return e.Err.Error()
}

func NewDomainError(code logcode.LogCode, message, localeKey string, err error) *DomainError {
	return &DomainError{
		Err:       err,
		Code:      code,
		Message:   message,
		LocaleKey: localeKey,
	}
}
