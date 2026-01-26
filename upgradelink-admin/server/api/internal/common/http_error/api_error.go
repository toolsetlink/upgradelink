package http_error

import (
	"net/http"
)

// NewApiError returns Api Error
func NewApiError(code int, msg string) error {
	return &LinkError{HttpCode: code, Code: code, Msg: msg}
}

// NewApiErrorWithoutMsg returns Api Error without message
func NewApiErrorWithoutMsg(code int) error {
	return &LinkError{HttpCode: code, Code: code, Msg: ""}
}

// NewApiInternalError returns Api Error with http internal error status code
func NewApiInternalError(msg string) error {
	return &LinkError{HttpCode: http.StatusInternalServerError, Code: http.StatusInternalServerError, Msg: msg}
}

// NewApiBadRequestError returns Api Error with http bad request status code
func NewApiBadRequestError(msg string) error {
	return &LinkError{HttpCode: http.StatusBadRequest, Code: http.StatusBadRequest, Msg: msg}
}

// NewApiUnauthorizedError returns Api Error with http unauthorized status code
func NewApiUnauthorizedError(msg string) error {
	return &LinkError{HttpCode: http.StatusUnauthorized, Code: http.StatusUnauthorized, Msg: msg}
}

// NewApiForbiddenError returns Api Error with http forbidden status code
func NewApiForbiddenError(msg string) error {
	return &LinkError{HttpCode: http.StatusForbidden, Code: http.StatusForbidden, Msg: msg}
}

// NewApiNotFoundError returns Api Error with http not found status code
func NewApiNotFoundError(msg string) error {
	return &LinkError{HttpCode: http.StatusForbidden, Code: http.StatusNotFound, Msg: msg}
}

// NewApiBadGatewayError returns Api Error with http bad gateway status code
func NewApiBadGatewayError(msg string) error {
	return &LinkError{HttpCode: http.StatusForbidden, Code: http.StatusBadGateway, Msg: msg}
}
