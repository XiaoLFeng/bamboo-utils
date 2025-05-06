/*
 * **********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * **********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * **********************************************************
 */

package bmodels

// ResponseDTO 一个通用的响应数据传输对象，用于 API 的返回值封装。
//
// 包含上下文、状态码、消息、时间戳、耗时信息及数据内容。支持泛型以适配不同类型数据。
//
// 参数
//   - T interface{}
//
// 结构体
//   - Context string 用于记录上下文信息
//   - Code uint 状态码
//   - Message string 消息返回
//   - Time int64 时间戳
//   - Overhead uint 耗时记录(毫秒)
//   - Data T 数据返回
type ResponseDTO[E interface{}] struct {
	Context  string `json:"context,omitempty" dc:"上下文"`
	Code     uint   `json:"code" dc:"状态码"`
	Message  string `json:"message" dc:"消息返回"`
	Time     int64  `json:"time" dc:"时间"`
	Overhead *int64 `json:"overhead,omitempty" dc:"耗时"`
	Data     E      `json:"data,omitempty" dc:"数据返回"`
}
