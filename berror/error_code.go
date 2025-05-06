package berror

// ErrorCode 表示错误码结构体，包含错误码、错误信息和扩展数据。
//
// # 结构
//   - Code 错误码，用于标识错误类型
//   - Message 错误信息，用于描述错误内容
//   - Data 错误相关数据，可扩展使用
type ErrorCode struct {
	Code    uint        `json:"code" description:"错误码"`
	Message string      `json:"message" description:"错误信息"`
	Data    interface{} `json:"data,omitempty" description:"错误数据"`
}

var (

	//
	// 400 Bad Request errors (40001-40099)
	//

	ErrBadRequest        = ErrorCode{Code: 40001, Message: "请求错误"}
	ErrInvalidParameters = ErrorCode{Code: 40002, Message: "无效的参数"}
	ErrMissingParameters = ErrorCode{Code: 40003, Message: "缺少必要参数"}
	ErrInvalidFormat     = ErrorCode{Code: 40004, Message: "格式无效"}
	ErrValidationFailed  = ErrorCode{Code: 40005, Message: "验证失败"}

	//
	// 401 Unauthorized errors (40101-40199)
	//

	ErrUnauthorized       = ErrorCode{Code: 40101, Message: "未授权"}
	ErrInvalidToken       = ErrorCode{Code: 40102, Message: "无效的令牌"}
	ErrTokenExpired       = ErrorCode{Code: 40103, Message: "令牌已过期"}
	ErrInvalidCredentials = ErrorCode{Code: 40104, Message: "无效的凭证"}
	ErrEmptyReferer       = ErrorCode{Code: 40105, Message: "请求头中缺少Referer"}

	//
	// 403 Forbidden errors (40301-40399)
	//

	ErrForbidden    = ErrorCode{Code: 40301, Message: "禁止访问"}
	ErrNoPermission = ErrorCode{Code: 40302, Message: "没有权限"}
	ErrAccessDenied = ErrorCode{Code: 40303, Message: "拒绝访问"}

	//
	// 404 Not Found errors (40401-40499)
	//

	ErrNotFound         = ErrorCode{Code: 40401, Message: "资源不存在"}
	ErrResourceNotFound = ErrorCode{Code: 40402, Message: "找不到指定资源"}
	ErrRouteNotFound    = ErrorCode{Code: 40403, Message: "路由不存在"}

	//
	// 429 Too Many Requests (42901-42999)
	//

	ErrTooManyRequests   = ErrorCode{Code: 42901, Message: "请求过多"}
	ErrRateLimitExceeded = ErrorCode{Code: 42902, Message: "超出频率限制"}

	//
	// 500 Internal Server errors (50001-50099)
	//

	ErrInternalServer    = ErrorCode{Code: 50001, Message: "服务器内部错误"}
	ErrDatabaseError     = ErrorCode{Code: 50002, Message: "数据库错误"}
	ErrCacheError        = ErrorCode{Code: 50003, Message: "缓存错误"}
	ErrUnexpectedError   = ErrorCode{Code: 50004, Message: "未预期的错误"}
	ErrDeveloperError    = ErrorCode{Code: 50005, Message: "开发者开发错误"}
	ErrUndefinitionError = ErrorCode{Code: 50006, Message: "未定义错误"}

	//
	// 503 Service Unavailable (50301-50399)
	//

	ErrServiceUnavailable = ErrorCode{Code: 50301, Message: "服务不可用"}
	ErrTimeout            = ErrorCode{Code: 50302, Message: "服务超时"}
	ErrOverloaded         = ErrorCode{Code: 50303, Message: "服务过载"}
)

// NewErrorCode 创建一个新的错误码实例。
//
// 参数:
//   - code: 错误码，用于标识错误类型
//   - message: 错误信息，用于描述错误内容
//   - data: 附加数据，提供额外的上下文信息
//
// 返回:
//   - 包含错误码、错误信息和扩展数据的 ErrorCode 实例
func NewErrorCode(code uint, message string, data interface{}) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// ErrorAddData 为错误代码添加关联数据，并返回更新后的错误代码。
//
// 参数:
//   - merr: 错误码结构体
//   - data: 要添加的关联数据
//
// 返回:
//   - 包含新增数据的错误码结构体
func ErrorAddData(err ErrorCode, data interface{}) *ErrorCode {
	err.Data = data
	return &err
}
