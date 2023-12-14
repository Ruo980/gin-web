package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

// 自定义声明结构体并内嵌 jwt.RegisteredClaims 结构体
type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

// 生成 token 函数:存储用户id与名称
func GenerateToken(id int, name string) (string, error) {
	// 获取 token 的密钥（KEY）
	var jwtKey = []byte(viper.GetString("jwt.secretKey"))
	// 创建 Payload：声明对象并填充数据
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// token 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 主题
			Subject: "Token",
		},
	}
	// 签发 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	// 加密：生成 JWT 并返回
	return token.SignedString(jwtKey)

}

// 解析 token 函数:解析token并返回用户id与名称，验证完整性
func ParseToken(tokenStr string) (JwtCustClaims, error) {
	// 获取 token 的密钥（KEY）
	var jwtKey = []byte(viper.GetString("jwt.secretKey"))
	// 解析 token
	claims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("Invalid token")
	}
	return claims, err
}

// 判断 token 是否非法
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
