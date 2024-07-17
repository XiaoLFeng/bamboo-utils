package butil

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"strings"
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

// StringToUUID
//
// # 字符串转UUID
//
// 用于将字符串转换为 UUID，返回转换后的 UUID. 转换后的 UUID 为 uuid.UUID 类型.
// 将字符串 UUID 转换为 UUID 类型.
//
// # 参数:
//   - getString 	传入的字符串(string)
//
// # 返回:
//   - getUUID 	转换后的 UUID(uuid.UUID)
func StringToUUID(getString string) (getUUID uuid.UUID) {
	parse, _ := uuid.Parse(getString)
	return parse
}

// TokenRemoveBearer
//
// # 移除Token中的Bearer
//
// 用于移除 Token 中的 "Bearer " 前缀，返回移除后, 前缀的 Token.
//
// # 参数:
//   - getToken 	带有 "Bearer " 前缀的 Token(string)
//
// # 返回:
//   - string 	去除 "Bearer " 前缀后的 Token
func TokenRemoveBearer(getToken string) string {
	return strings.Replace(getToken, "Bearer ", "", 1)
}

// PasswordEncode
//
// # 密码加密
//
// 用于对密码进行加密，返回加密后的密码. 加密规则遵循 bcrypt 加密算法. 返回的密码为 string 类型.
// 加密方式为: base64 -> md5 -> bcrypt.
//
// # 参数:
//   - getPassword 	需要加密的密码(string)
//
// # 返回:
//   - string 	加密后的密码
func PasswordEncode(getPassword string) string {
	// 首先对密码进行双重加密
	getBase64 := base64.StdEncoding.EncodeToString([]byte(getPassword))
	getMd5Password := md5.New().Sum([]byte(getBase64))
	encodePassword, err := bcrypt.GenerateFromPassword(getMd5Password, bcrypt.DefaultCost)
	if err != nil {
		return ""
	} else {
		return string(encodePassword)
	}
}

// PasswordVerify
//
// # 密码验证
//
// 用于验证密码是否正确，返回验证结果. 验证规则遵循 bcrypt 加密算法. 返回的结果为 bool 类型.
// 验证方式为: base64 -> md5 -> bcrypt.
//
// # 参数:
//   - getPassword 		原始密码(string)
//   - getHashPassword 	需要验证的密码(string)
//
// # 返回:
//   - bool 	验证结果
func PasswordVerify(getPassword, getHashPassword string) bool {
	// 首先对密码进行双重加密
	getBase64 := base64.StdEncoding.EncodeToString([]byte(getPassword))
	getMd5Password := md5.New().Sum([]byte(getBase64))
	err := bcrypt.CompareHashAndPassword([]byte(getHashPassword), getMd5Password)
	return err == nil
}

// RandomString
//
// # 生成随机字符串
//
// 用于生成一个指定长度的随机字符串，返回生成的随机字符串. 生成的字符串为 string 类型.
//
// # 参数:
//   - length 	生成的字符串长度(int)
//
// # 返回:
//   - string 	生成的随机字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[randIndex.Int64()]
	}
	return string(result)
}

// GetVisitorProtocol
//
// # 获取访问者协议
//
// 用于获取访问者使用的协议类型，返回协议类型. 返回的协议类型为 string 类型. 如果是 SSL 协议则返回 "SSL"，否则返回 "TLS".
// 即若请求中访问协议为 HTTPS 则返回 "SSL"，否则返回 "TLS".
// 若 context.Context 为空则返回空字符串.
//
// # 参数:
//   - ctx 	请求的上下文(context.Context)
//
// # 返回:
//   - string 	协议类型
func GetVisitorProtocol(ctx context.Context) string {
	getRequest := ghttp.RequestFromCtx(ctx)
	if getRequest != nil {
		if getRequest.TLS != nil {
			return "SSL"
		} else {
			return "TLS"
		}
	} else {
		return ""
	}
}
