package http_error

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/trace"
)

type LinkError struct {
	HttpCode int    `json:"http_code"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}

// Error 用于错误处理
func (e *LinkError) Error() string {
	return fmt.Sprintf("[%d]: %s", e.Code, e.Msg)
}

type ErrResp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Docs    string `json:"docs"`
	TraceId string `json:"traceId"`
}

// ErrorHandler 定义异常返回格式化信息
//
//	@desc	定义异常返回格式化信息
//	@param	context.Context
//	@param	error
//	@return	int
//	@return	any
func ErrorHandler(ctx context.Context, err error) (int, any) {
	traceID := trace.TraceIDFromContext(ctx)
	fmt.Println("ErrorHandler:", ErrorHandler)
	fmt.Println("ErrorHandler err:", err)
	if err == nil {
		return http.StatusOK, &ErrResp{Code: 0, Msg: "ok", TraceId: traceID}
	}

	var e *LinkError
	if errors.As(err, &e) {
		return e.HttpCode, &ErrResp{Code: e.Code, Msg: e.Msg, TraceId: traceID}
	}

	return http.StatusOK, &ErrResp{Code: 0, Msg: "ok", TraceId: traceID}
}
