package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       uint
	Username string
}

type Claims struct {
	jwt.RegisteredClaims
	User User
}

// GenerateJwt 生成token
func GenerateJwt(claims Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseJwt 解密token
func ParseJwt(token, secret string) (*Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}
