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
	"fmt"
	"strings"
)

// stringBuild 构建格式化的日志消息字符串。
//
// 参数:
//   - function: 功能名，用于标识日志来源
//   - message: 消息内容，描述日志的具体信息
//
// 返回:
//   - 拼接后的字符串，包含控制器、功能名及消息内容
func stringBuild(stringBuilder *strings.Builder, function string, message string, v ...interface{}) string {
	stringBuilder.WriteString(" ")
	stringBuilder.WriteString("<")
	stringBuilder.WriteString(function)
	stringBuilder.WriteString("> ")
	if len(v) > 0 {
		stringBuilder.WriteString(fmt.Sprintf(message, v...))
	} else {
		stringBuilder.WriteString(message)
	}
	return fmt.Sprintf(stringBuilder.String())
}
