package http_error

import (
	"net/http"
)

// NewCodeError returns Code Error
func NewCodeError(code int, msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: code, Msg: msg}
}

// NewCodeErrorWithoutMsg returns Code Error without message
func NewCodeErrorWithoutMsg(code int) error {
	return &LinkError{HttpCode: http.StatusOK, Code: code, Msg: ""}
}

// NewCodeInternalError returns Code Error with http internal error status code
func NewCodeInternalError(msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: http.StatusInternalServerError, Msg: msg}
}

// NewCodeBadRequestError returns Code Error with http bad request status code
func NewCodeBadRequestError(msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: http.StatusBadRequest, Msg: msg}
}

// NewCodeUnauthorizedError returns Code Error with http unauthorized status code
func NewCodeUnauthorizedError(msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: http.StatusUnauthorized, Msg: msg}
}

// NewCodeForbiddenError returns Code Error with http forbidden status code
func NewCodeForbiddenError(msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: http.StatusForbidden, Msg: msg}
}

// NewCodeNotFoundError returns Code Error with http not found status code
func NewCodeNotFoundError(msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: http.StatusNotFound, Msg: msg}
}

// NewCodeBadGatewayError returns Code Error with http bad gateway status code
func NewCodeBadGatewayError(msg string) error {
	return &LinkError{HttpCode: http.StatusOK, Code: http.StatusBadGateway, Msg: msg}
}
