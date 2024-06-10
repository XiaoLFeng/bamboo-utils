package test

import (
	"context"
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

// TestPasswordEncode
//
// # 测试密码加密
//
// 用于测试密码加密的功能是否正常.
func TestPasswordEncode(t *testing.T) {
	password := "admin-admin"
	encodePassword := butil.PasswordEncode(password)
	t.Logf("Password: %s", encodePassword)
	// 对密码进行校验
	if butil.PasswordVerify("admin-admin", encodePassword) {
		t.Logf("Password Verify: %s", "Success")
	} else {
		t.Logf("Password Verify: %s", "Failed")
	}
	if butil.PasswordVerify("admin", encodePassword) {
		t.Logf("Password Verify: %s", "Success")
	} else {
		t.Logf("Password Verify: %s", "Failed")
	}
}

func TestVisitorProtocol(t *testing.T) {
	butil.GetVisitorProtocol(context.Background())
}
