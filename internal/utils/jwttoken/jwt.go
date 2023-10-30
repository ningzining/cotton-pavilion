package jwttoken

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user-center/internal/consts"
)

type User struct {
	Id       uint64
	Username string
}

type CustomerClaims struct {
	jwt.RegisteredClaims
	User User
}

// Generate 生成token
func Generate(user User) (string, error) {
	claims := CustomerClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   consts.SystemName,
			Subject:  consts.JwtSubject,
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(consts.JwtSecretKey))
	return signedString, err
}

// Parse 解密token
func Parse(token string) (*CustomerClaims, error) {
	var claims CustomerClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}
