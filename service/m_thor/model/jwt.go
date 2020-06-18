/*
	鉴权主体
	author：Kyda
	date:06.18.2020
*/

package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey      = "this is a secretKey"            // 密钥
	ExpireDuration = time.Second * time.Duration(300) // 过期时间长度
)


// ParseToken 解析Token
func ParseToken(tokenStr string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return
}

// CheckToken 生成Token
func CheckToken(tokenStr string) bool {
	token, err := ParseToken(tokenStr)
	// 过期、无法识别等其他错误
	if err != nil || !token.Valid {
		return false
	}
	return true
}
