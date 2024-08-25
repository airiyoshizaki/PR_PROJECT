package apierr

import (
	"errors"
	"net/http"

	"github.com/route-test/constant/logcode"
	dc "github.com/route-test/domain/common"
)

type APIError struct {
	Code             string `json:"code" validate:"required"`
	StatusCode       int    `json:"status_code" validate:"required"`
	MessageLocaleKey string `json:"message_locale_key" validate:"required"`
}

// New 独自のエラーを生成する。APIのエラーレスポンスとなるため、詳細情報はエラーメッセージとしては含めない。
func New(err error, code logcode.LogCode, msgLocaleKey string) *APIError {
	e := APIError{}
	e.MessageLocaleKey = msgLocaleKey
	e.Code = code.Value

	switch {
	case errors.Is(err, dc.ErrBadRequest):
		e.StatusCode = http.StatusBadRequest

	case errors.Is(err, dc.ErrUnauthorized):
		e.StatusCode = http.StatusUnauthorized

	case errors.Is(err, dc.ErrForbidden):
		e.StatusCode = http.StatusForbidden

	case errors.Is(err, dc.ErrNotFound):
		e.StatusCode = http.StatusNotFound

	case errors.Is(err, dc.ErrServiceUnavailable):
		e.StatusCode = http.StatusServiceUnavailable

	default:
		e.StatusCode = http.StatusInternalServerError
	}

	return &e
}
