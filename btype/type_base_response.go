package btype

// BaseResponse
//
// # 基本数据返回
//
// 基本数据结构，进行数据信息的返回
//
// # 参数
//   - Output 			英文输出(string)
//   - Code 			状态码(int)
//   - Message 			中文描述(string)
//   - ErrorMessage		错误信息(string?)
//   - Data				数据(interface?)
type BaseResponse struct {
	Output       string      `json:"output" dc:"英文输出"`
	Code         int         `json:"code" dc:"状态码"`
	Message      string      `json:"message" dc:"中文描述"`
	ErrorMessage string      `json:"error_message,omitempty" dc:"错误信息"`
	Data         interface{} `json:"data,omitempty" dc:"数据"`
}
