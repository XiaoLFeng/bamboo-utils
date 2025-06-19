/*
 * **********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * **********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * **********************************************************
 */

package test

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	debugInterface(t, "你好 %v:%v:%v:%v:%v", "Hello, World!", 42, 3.14, true, []int{1, 2, 3})

	debugInterface(t, "你好")
}

func debugInterface(t *testing.T, format string, value ...interface{}) {
	for _, v := range value {
		switch v.(type) {
		case string:
			t.Logf("string: %s", v)
		case int:
			t.Logf("int: %d", v)
		case float64:
			t.Logf("float64: %f", v)
		default:
			t.Logf("unknown type: %v", v)
		}
	}
	t.Logf("%v", len(value))
	if len(value) > 0 {
		t.Log(fmt.Sprintf(format, value...))
	} else {
		t.Log(format)
	}
}
