/*
 * **********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * **********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * **********************************************************
 */

package private

import (
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GErrorTransform 将 gerror 错误对象转换为 berror 错误码。
//
// 该函数将传入的错误对象转换为 berror.ErrorCode 错误码。请注意，建议在 bmiddle.BambooHandlerResponse 中使用。
// 尽量不自行进行调用，否则容易出现不可预测情况。
//
// 参数:
//   - err: 传入的错误对象，支持包含 Code 方法的错误类型。
//
// 返回:
//   - 转换后的 berror.ErrorCode 错误码。
//
// 错误:
//   - 如果输入的错误对象未包含有效的 Code 信息，返回 berror.ErrUndefinitionError。
func GErrorTransform(errCode *gerror.Error) berror.ErrorCode {

	switch errCode.Code() {
	case gcode.CodeInternalError:
		return berror.ErrInternalServer
	case gcode.CodeValidationFailed:
		return *berror.ErrorAddData(berror.ErrValidationFailed, errCode.Error())
	case gcode.CodeDbOperationError:
		return berror.ErrDatabaseError
	case gcode.CodeInvalidParameter:
		return berror.ErrInvalidParameters
	case gcode.CodeMissingParameter:
		return berror.ErrMissingParameters
	case gcode.CodeInvalidOperation:
		return berror.ErrInvalidOperation
	case gcode.CodeInvalidConfiguration:
		return berror.ErrInvalidConfiguration
	case gcode.CodeMissingConfiguration:
		return berror.ErrMissingConfiguration
	case gcode.CodeNotImplemented:
		return berror.ErrNotImplemented
	case gcode.CodeNotSupported:
		return berror.ErrNotSupported
	case gcode.CodeOperationFailed:
		return berror.ErrUnexpectedError
	case gcode.CodeNotAuthorized:
		return berror.ErrUnauthorized
	case gcode.CodeSecurityReason:
		return berror.ErrForbidden
	case gcode.CodeServerBusy:
		return berror.ErrOverloaded
	case gcode.CodeUnknown:
		return berror.ErrUnexpectedError
	case gcode.CodeNotFound:
		return berror.ErrNotFound
	case gcode.CodeInvalidRequest:
		return berror.ErrBadRequest
	case gcode.CodeNecessaryPackageNotImport:
		return berror.ErrInternalDependency
	case gcode.CodeInternalPanic:
		return berror.ErrInternalPanic
	case gcode.CodeBusinessValidationFailed:
		return berror.ErrBusinessValidationFailed

	// 处理 CodeNil 和 CodeOK 的情况
	case gcode.CodeNil:
		return berror.ErrUndefinitionError
	case gcode.CodeOK:
		return berror.ErrUndefinitionError

	default:
		// 对于任何未明确列出的 gcode.Code，返回未定义错误
		// 这也包括 err.Code() 返回 nil 的情况 (已在函数开始处处理)
		return berror.ErrUndefinitionError
	}
}
