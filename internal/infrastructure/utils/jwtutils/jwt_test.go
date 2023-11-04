package jwtutils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestGenerateJwt(t *testing.T) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "issuer",
			Subject:  "subject",
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		User: User{
			Id:       1,
			Username: "cotton",
		},
	}
	secret := "cotton"
	token, err := GenerateJwt(claims, secret)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", token)
}

func TestParseJwt(t *testing.T) {
	secret := "cotton"
	token, err := ParseJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsInN1YiI6ImF1dGgiLCJpYXQiOjE2OTg4NDU1MjYsIlVzZXIiOnsiSWQiOjEsIlVzZXJuYW1lIjoiY290dG9uIn19.AfxrOVVqNh4qCchJ6oPOHiKKulEK0l6UgXSCWlzf_0A", secret)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%+v\n", token.User)
}
