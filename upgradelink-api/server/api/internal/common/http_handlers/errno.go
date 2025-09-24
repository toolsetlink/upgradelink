// Package http_handlers
//	@Desc: 通用错误码处理方式

package http_handlers

// OK
//
//	@状态码: 200
//	@状态含义: Normal
//	@状态原因: 无异常
//	@错误码
const (
	OK = 0 // 正常

	NoContent = 204001 // 无内容
)

// 服务端错误码

// ErrInternalServerError
//
//	@状态码: 500
//	@状态含义: Internal server error
//	@状态原因: 客户端请求有效, 服务器处理时发生了意外!
//	@错误码
const (
	ErrInternalServerError = 500001 // 服务器内部错误
	ErrServerCheckError    = 500002 // 数据效验错误
)

// ErrNotImplemented
//
//	@状态码: 501
//	@状态含义: Not Implemented
//	@状态原因: 服务器不支持请求的功能，无法完成请求
//	@错误码
const (
	ErrNotImplemented = 501001 // 不支持的请求
)

// ErrBadGateway
//
//	@状态码: 502
//	@状态含义: Bad Gateway
//	@状态原因: 作为网关或者代理工作的服务器尝试执行请求时，从远程服务器接收到了一个无效的响应
//	@错误码
const (
	ErrBadGateway = 502001 // 无效的响应
)

// ErrSvcUnavailable
//
//	@状态码: 503
//	@状态含义: Service unavailable
//	@状态原因: 服务器无法处理请求, 一般用于网站维护状态!
//	@错误码
const (
	ErrSvcUnavailable = 503001 // 服务不可达
)

// 2.客户端端错误码

// ErrBadReq
//
//	@状态码: 400
//	@状态含义: Bad request
//	@状态原因: 服务器不理解客户端的请求, 未做任何处理!
//	@错误码
const (
	ErrBadReq       = 400001 // 非法请求
	ErrParamMiss    = 400002 // 参数缺失
	ErrParamInvalid = 400003 // 参数非法
	ErrHeadInvalid  = 400004 // 报头非法
	ErrBodyInvalid  = 400005 // 报体非法
)

// ErrAuth
//
//	@状态码: 401
//	@状态含义: Unauthorized
//	@状态原因: 用户未提供身份验证凭据, 或者没有通过身份验证!
//	@错误码
const (
	ErrAuth = 401001 // 鉴权失败
	ErrSign = 401002 // 签名验证失败
)

// ErrForbidden
//
//	@状态码: 403
//	@状态含义: Forbidden
//	@状态原因: 用户通过了身份验证, 但是不具有访问资源所需的权限!
//	@错误码
const (
	ErrForbidden = 403001 // 访问受限
	ErrUsrFrozen = 403002 // 用户被冻结
)

// ErrNotFound
//
//	@状态码: 404
//	@状态含义: Not found
//	@状态原因: 所请求的资源不存在, 或不可用!
//	@错误码
const (
	ErrNotFound        = 404001 // 资源不存在
	ErrRecordNotExist  = 404002 // 查询记录不存在
	ErrHandlerNotExist = 404003 // 处理方式不存在
	ErrRequestNotFound = 404004 // 请求不存在
)

// ErrMethodNotAllowed
//
//	@状态码: 405
//	@状态含义: Method not allowed
//	@状态原因: 用户已经通过身份验证, 但是所用的 HTTP 方法不在他的权限之内!
//	@错误码
const (
	ErrMethodNotAllowed = 405001 // 无本 HTTP 访问权限
)

// ErrTimeOut
//
//	@状态码: 408
//	@状态含义: timeout
//	@状态原因: 请求超时!
//	@错误码
const (
	ErrTimeOutChat = 408001 // 请求 gpt 超时
)

// ErrConflict
//
//	@状态码: 409
//	@状态含义: Conflict
//	@状态原因: 由于和被请求的资源的当前状态之间存在冲突，请求无法完成
//	@错误码
const (
	ErrConflict = 409001 // 资源冲突
)

// ErrGone
//
//	@状态码: 410
//	@状态含义: Gone
//	@状态原因: 所请求的资源已从这个地址转移, 不再可用!
//	@错误码
const (
	ErrGone = 410001 // 资源不再可用
)

// ErrUnsupportedMediaType
//
//	@状态码: 415
//	@状态含义: Unsupported media type
//	@状态原因: 客户端要求的返回格式不支持. 比如: API 只能返回 JSON 格式, 但是客户端要求返回 XML 格式!
//	@错误码
const (
	ErrUnsupportedMediaType = 415001 // 不支持的返回格式
)

// ErrUnprocessableEntity
//
//	@状态码: 422
//	@状态含义: Unprocessable entity
//	@状态原因: 客户端上传的附件无法处理, 导致请求失败!
//	@错误码
const (
	ErrUnprocessableEntity = 422001 // 不支持的返回格式
)

// ErrTooManyReq
//
//	@状态码: 429
//	@状态含义: Too many requests
//	@状态原因: 客户端的请求次数超过限额!
//	@错误码
const (
	ErrStrategyTooManyReq = 429001 // 策略请求次数超过限制
)

const (
	ErrServerCheckError001 = 500002001 // 请求url应用不存在
)
