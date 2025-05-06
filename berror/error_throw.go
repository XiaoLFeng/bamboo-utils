package berror

// Error 表示一个通用错误接口，用于描述错误的详细信息及编码。
//
// 该接口提供以下功能:
//   - 返回错误码对象 (GetErrorCode)。
//   - 返回错误信息字符串 (Error)。
//   - 返回更详细的错误消息 (GetErrorMessage)。
type Error interface {
	GetErrorCode() ErrorCode
	Error() string
	GetErrorMessage() string
}

// GetErrorCode 返回当前的错误码对象。
//
// 返回:
//   - 当前的 ErrorCode 结构体副本。
func (e *ErrorCode) GetErrorCode() ErrorCode {
	return *e
}

// Error 返回错误的信息描述。
//
// 返回:
//   - 错误的描述信息字符串。
func (e *ErrorCode) Error() string {
	return e.Message
}

// GetErrorMessage 返回错误信息字符串。
func (e *ErrorCode) GetErrorMessage() string {
	return e.Message
}
