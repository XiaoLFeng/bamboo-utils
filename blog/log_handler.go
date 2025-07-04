/*
 * **********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * **********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * **********************************************************
 */

package blog

import (
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bconsts"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

// HandlerInfo 输出中间件相关日志。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的消息内容。
//   - v: 可选参数，用于格式化消息内容。
func HandlerInfo(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.HANDLER)
	g.Log().Info(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// HandlerDebug 输出工具相关的调试日志。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的消息内容。
//   - v: 可选参数，用于格式化消息内容。
func HandlerDebug(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.HANDLER)
	g.Log().Debug(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// HandlerError 输出工具相关的错误日志。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的错误消息内容。
//   - v: 可选参数，用于格式化错误消息内容。
func HandlerError(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.HANDLER)
	g.Log().Error(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// HandlerNotice 记录通知级别的日志消息。
//
// 参数:
//   - ctx: 上下文对象，传递请求级别的信息
//   - function: 功能名，用于标识日志来源
//   - format: 消息格式字符串
//   - v: 格式化内容的参数列表
func HandlerNotice(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.HANDLER)
	g.Log().Notice(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// HandlerPanic 输出工具相关的严重错误日志并触发panic。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的严重错误消息内容。
//   - v: 可选参数，用于格式化严重错误消息内容。
func HandlerPanic(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.HANDLER)
	g.Log().Panic(ctx, stringBuild(&stringBuilder, function, format, v...))
}
