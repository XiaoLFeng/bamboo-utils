package berror

import "github.com/bamboo-services/bamboo-utils/bcode"

// IError
//
// # 基本报错
type IError interface {
	getErrorInfo() error
	getECode() *bcode.BCode
}

// getErrorInfo
//
// # 获取报错信息
func (e *XError) getErrorInfo() error {
	return e.error
}

// getECode
//
// # 获取报错代码
func (e *XError) getECode() *bcode.BCode {
	return &e.code
}
