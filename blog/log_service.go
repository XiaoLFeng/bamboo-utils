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

// ServiceInfo 输出服务层相关的信息日志。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的消息内容。
//   - v: 可变参数，用于格式化消息内容。
func ServiceInfo(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.SERVICE)
	g.Log().Info(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// ServiceDebug 输出服务层相关的调试日志。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的消息内容。
//   - v: 可变参数，用于格式化消息内容。
func ServiceDebug(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.SERVICE)
	g.Log().Debug(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// ServiceError 输出服务层相关的错误日志。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的错误消息内容。
//   - v: 可变参数，用于格式化错误消息内容。
func ServiceError(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.SERVICE)
	g.Log().Error(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// ServiceNotice 输出服务级别的通知日志。
//
// 参数:
//   - ctx: 上下文信息，用于日志关联上下文数据
//   - function: 功能名，用于标识日志来源
//   - format: 日志内容的格式化字符串
//   - v: 可变参数，用于格式化日志内容
func ServiceNotice(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.SERVICE)
	g.Log().Notice(ctx, stringBuild(&stringBuilder, function, format, v...))
}

// ServicePanic 输出服务层相关的严重错误日志并触发panic。
//
// 参数:
//   - ctx: 上下文对象，用于日志记录。
//   - function: 功能名，用于标识日志来源。
//   - format: 要记录的严重错误消息内容。
//   - v: 可变参数，用于格式化严重错误消息内容。
func ServicePanic(ctx context.Context, function string, format string, v ...interface{}) {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(bconsts.SERVICE)
	g.Log().Panic(ctx, stringBuild(&stringBuilder, function, format, v...))
}
