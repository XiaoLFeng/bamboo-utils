package butil

import (
	"context"
	"crypto/rand"
	"github.com/gogf/gf/v2/net/ghttp"
	"math/big"
	"strings"
)

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
