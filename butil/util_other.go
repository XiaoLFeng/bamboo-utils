package butil

import (
	"context"
	"crypto/rand"
	"github.com/gogf/gf/v2/net/ghttp"
	"math/big"
	"strings"
)

// TokenRemoveBearer 移除字符串中的 "Bearer " 前缀并返回处理后的结果。
//
// 参数:
//   - getToken 	要处理的字符串(string)
//
// 返回:
//   - string 	处理后的结果
func TokenRemoveBearer(getToken string) string {
	return strings.Replace(getToken, "Bearer ", "", 1)
}

// RandomString 生成指定长度的随机字符串，由字母和数字组成。
//
// 参数:
//   - length: 要生成的字符串长度
//
// 返回:
//   - 包含随机字符的字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[randIndex.Int64()]
	}
	return string(result)
}

// GetVisitorProtocol 根据上下文获取访问协议标识。
//
// 参数:
//   - ctx: 请求的上下文对象。
//
// 返回:
//   - 返回访问协议的标识，如 "SSL" 或 "TLS"。若请求对象为空，则返回空字符串。
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
