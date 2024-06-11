package bcode

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

// BCode
//
// # 错误码接口
//
// 错误码接口, 用于返回错误码信息.
//
// # 方法:
//   - Code		错误码(int)
//   - Message		错误信息(string)
//   - Output		报错英文规范说明(string)
type BCode interface {
	// Code - 返回一个错误码的具体信息
	Code() int
	// Message - 返回错误码所对应的中文解释大致错误
	Message() string
	// Output - 返回一个英文规范错误说明
	Output() string
}

// ==================================================
// 常见错误代码定义。
// 保留了内部错误代码：代码 100 <= HERE <= 200。
// ==================================================

var (
	Success                   = LocalCode{output: "Success", code: 200, message: "操作成功"}
	ServerInternalError       = LocalCode{output: "ServerInternalError", code: 101, message: "服务器内部错误"}
	RequestParameterIncorrect = LocalCode{output: "RequestParameterIncorrect", code: 103, message: "请求参数不正确"}
	RequestBodyIncorrect      = LocalCode{output: "RequestBodyIncorrect", code: 104, message: "请求体内容不正确"}
	RequestPathIncorrect      = LocalCode{output: "RequestPathIncorrect", code: 105, message: "请求路径不正确"}
	RequestHeaderIncorrect    = LocalCode{output: "RequestHeaderIncorrect", code: 106, message: "请求头不正确"}
	RequestMethodIncorrect    = LocalCode{output: "RequestMethodIncorrect", code: 107, message: "请求方法不正确"}
	RequestContentTypeError   = LocalCode{output: "RequestContentTypeError", code: 108, message: "请求内容类型错误"}
	RequestContentLengthError = LocalCode{output: "RequestContentLengthError", code: 109, message: "请求内容长度错误"}
	UnknownError              = LocalCode{output: "UnknownError", code: 110, message: "未知错误"}
	PageNotFound              = LocalCode{output: "PageNotFound", code: 111, message: "页面不存在"}
	StatusForbidden           = LocalCode{output: "StatusForbidden", code: 112, message: "禁止访问"}
	AlreadyExists             = LocalCode{output: "AlreadyExists", code: 113, message: "内容已存在"}
	NotExist                  = LocalCode{output: "NotExist", code: 114, message: "内容不存在"}
	Expired                   = LocalCode{output: "Expired", code: 115, message: "内容已过期"}
	OperationFailed           = LocalCode{output: "OperationFailed", code: 116, message: "操作失败"}
	OperationTimeout          = LocalCode{output: "OperationTimeout", code: 117, message: "操作超时"}
	OperationLimit            = LocalCode{output: "OperationLimit", code: 118, message: "操作限制"}
	OperationNotSupport       = LocalCode{output: "OperationNotSupport", code: 119, message: "操作不支持"}
	OperationNotAllow         = LocalCode{output: "OperationNotAllow", code: 120, message: "操作不允许"}
	OperationNotImplement     = LocalCode{output: "OperationNotImplement", code: 121, message: "操作未实现"}
	OperationNotPermission    = LocalCode{output: "OperationNotPermission", code: 122, message: "操作无权限"}
	OperationNotAvailable     = LocalCode{output: "OperationNotAvailable", code: 123, message: "操作不可用"}
	VerifyFailed              = LocalCode{output: "VerifyFailed", code: 124, message: "验证失败"}
	VerifyTimeout             = LocalCode{output: "VerifyTimeout", code: 125, message: "验证超时"}
	VerifyLimit               = LocalCode{output: "VerifyLimit", code: 126, message: "验证限制"}
	VerifyNotSupport          = LocalCode{output: "VerifyNotSupport", code: 127, message: "验证不支持"}
	VerifyNotAllow            = LocalCode{output: "VerifyNotAllow", code: 128, message: "验证不允许"}
)

// BaseLocalCode
//
// # 基础错误码
//
// 用于创建一个基础的错误码, 用于返回错误码信息.
//
// # 参数
//   - output		报错英文规范说明(string)
//   - code		错误码(int)
//   - message		错误信息(string)
//
// # 返回
//   - BCode		本地错误码
func BaseLocalCode(output string, code int, message string) BCode {
	return LocalCode{output: output, code: code, message: message}
}

// GcodeToLocalCode
//
// # Gcode转换为本地错误码
//
// 用于将gcode转换为本地错误码, 用于返回错误码信息. Gcode 为 gogf/gf/v2/errors/gcode 包中的错误码.
//
// # 参数
//   - gcode		gcode错误码
//
// # 返回
//   - BCode		本地错误码
func GcodeToLocalCode(gcode gcode.Code) BCode {
	return BaseLocalCode("GCodeError", gcode.Code(), gcode.Message())
}
