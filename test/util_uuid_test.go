package test

import (
	"github.com/bamboo-services/bamboo-utils/butil"
	"testing"
)

// TestMakeUUID
//
// # 测试生成UUID
//
// 用于测试生成UUID的功能是否正常.
func TestMakeUUID(t *testing.T) {
	uuid := butil.GenerateRandUUID()
	t.Logf("RandUUID: %s", uuid)

	makeStringUUIDFirst := butil.MakeUUIDByString("xiao_lfeng")
	t.Logf("MakeUUID: %s", makeStringUUIDFirst)

	makeStringUUIDSecond := butil.MakeUUIDByString("xiao_lfeng")
	t.Logf("MakeUUID: %s", makeStringUUIDSecond)
}
