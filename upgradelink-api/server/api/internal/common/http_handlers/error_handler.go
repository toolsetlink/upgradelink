// Package http_handlers 乐视通用错误码处理方式
//
//	@Desc: 乐视通用错误码处理方式
package http_handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/zeromicro/go-zero/core/trace"
)

// ErrResp 报错回包结构体
type ErrResp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Docs    string `json:"docs"`
	TraceId string `json:"traceId"`
}

// 异常类型正则解析器
var (
	linkErrReg        = regexp.MustCompile(`^\[\d{6}]: .+`)
	paramMissErrReg   = regexp.MustCompile(`^field .+ is not set$`)
	paramTypeErrReg   = regexp.MustCompile(`^type mismatch for field .+`)
	paramRageErrReg   = regexp.MustCompile(`^wrong number range setting$`)
	paramOptionErrReg = regexp.MustCompile(`^error: value .+ is not defined in options .+`)
)

// ErrorHandler 定义异常返回格式化信息
//
//	@desc	定义异常返回格式化信息
//	@param	context.Context
//	@param	error
//	@return	int
//	@return	any
func ErrorHandler(ctx context.Context, err error) (int, any) {
	statusCode, code, msg, docs, _ := LinkErr2Code(err)
	traceID := trace.TraceIDFromContext(ctx)
	return statusCode, &ErrResp{Code: code, Msg: msg, Docs: docs, TraceId: traceID}
}

// ErrMsg2Code 异常消息转 Code
//
//	@desc    异常消息转 Code
//	@param
//	@return
func ErrMsg2Code(m string) (httpCode int, code int, msg, docs, traceId string) {
	if linkErrReg.MatchString(m) {
		code, _ = strconv.Atoi(m[1:7])
		return code / 1000, code, m[10:], docs, ""
	}

	if paramMissErrReg.MatchString(m) {
		return ErrParamMiss / 1000, ErrParamMiss, m, docs, ""
	}

	if paramTypeErrReg.MatchString(m) ||
		paramRageErrReg.MatchString(m) || paramOptionErrReg.MatchString(m) {
		return ErrParamInvalid / 1000, ErrParamInvalid, m, docs, ""
	}

	return http.StatusInternalServerError, ErrInternalServerError, m, docs, ""
}

// //////////////////////////////////////////////////////////////////////////////
// 定义 LeError

type linkErr struct {
	statusCode int
	code       int
	msg        string
	docs       string
	traceId    string
}

// Error 用于错误处理
func (e *linkErr) Error() string {
	return fmt.Sprintf("[%d]: %s", e.code, e.msg)
}

// NewLinkErr 创建 err
func NewLinkErr(ctx context.Context, code int, msg string, docs string) error {
	traceID := trace.TraceIDFromContext(ctx)
	codeStr := strconv.Itoa(code)
	if len(codeStr) == 9 {
		return &linkErr{statusCode: code / 1000000, code: code, msg: msg, docs: docs, traceId: traceID}
	}

	return &linkErr{statusCode: code / 1000, code: code, msg: msg, docs: docs, traceId: traceID}
}

// LinkErr2Code 将 Error 格式为 code
//
//	@desc    将 Error 格式为 code
//	@param   err        error
//	@return  statusCode int "http code"
//	@return  code       int "业务码"
func LinkErr2Code(err error) (statusCode, code int, msg, docs, traceId string) {
	if err == nil {
		return http.StatusOK, OK, "ok", "", ""
	}

	var e *linkErr
	if errors.As(err, &e) {
		return e.statusCode, e.code, e.msg, e.docs, e.traceId
	}

	return ErrMsg2Code(err.Error())
}

// IsLeMsg 判断是否为自定义异常
//
//	@desc    判断是否为自定义异常
//	@param
//	@return
func IsLeMsg(m string) bool {
	if len(m) < 11 {
		return false
	}
	_, terr := strconv.Atoi(m[1:7])
	if nil != terr {
		return false
	}
	return true
}
