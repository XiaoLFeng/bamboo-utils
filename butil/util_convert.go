/*
 * **********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * **********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * **********************************************************
 */

package butil

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// StructToMap 将结构体转换为 `map[string]interface{}`
//
// 结构体转为 Map 后，会自动将 string("") 变为 nil 形式
//
// 参数：
//   - v: 结构体
//
// 返回：
//   - map[string]interface{}
func StructToMap(v interface{}) map[string]interface{} {
	getMap := gconv.Map(v)
	for k, v := range getMap {
		if str, ok := v.(string); ok && str == "" {
			getMap[k] = nil
		}
	}
	return getMap
}

// MapToStruct 将 map 数据转换为指定结构体类型。
//
// 参数:
//   - value: 包含数据的 map[string]interface{}
//   - target: 要填充的目标结构体实例
//
// 返回:
//   - 填充后的目标结构体实例
//   - 转换过程中可能产生的错误
func MapToStruct[E interface{}](value map[string]interface{}, target E) (E, error) {
	for k, v := range value {
		if v == "null" {
			value[k] = nil
		}
	}
	var newTarget E
	err := gconv.Struct(value, &newTarget)
	return newTarget, err
}
