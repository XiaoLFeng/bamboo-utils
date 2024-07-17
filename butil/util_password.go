package butil

import (
	"crypto/md5"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

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
