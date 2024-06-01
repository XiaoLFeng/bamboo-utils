package butil

import (
	"crypto/md5"
	"github.com/google/uuid"
)

// GenerateRandUUID
//
// # 生成随机UUID
//
// 用于生成一个随机 UUID，包装使用了 github.com/google/uuid 库. 生成的规则遵循 Google UUID V7 版本.
// 生成的 UUID 为 uuid.UUID 类型.
//
// # 返回:
//   - getUUID 		生成随机的 UUID(uuid.UUID)
func GenerateRandUUID() (getUUID uuid.UUID) {
	get, _ := uuid.NewV7()
	return get
}

// MakeUUIDByString
//
// # 通过字符串生成UUID
//
// 用于通过字符串生成一个 UUID，生成的 UUID 是一个固定的值，不会随机生成.根据输入参数的内容决定返回的 UUID.
//
// # 参数:
//   - getString 	传入的字符串(string)
//
// # 返回:
//   - getUUID 	生成的 UUID(uuid.UUID)
func MakeUUIDByString(getString string) (getUUID uuid.UUID) {
	return md5.Sum([]byte(getString))
}
