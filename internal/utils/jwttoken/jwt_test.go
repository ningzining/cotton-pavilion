package jwttoken

import (
	"fmt"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	user := User{
		Id:       1,
		Username: "cotton",
	}
	token, err := GenerateJwtToken(user)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", token)
}

func TestParseJwtToken(t *testing.T) {
	token, err := ParseJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsInN1YiI6InVzZXItY2VudGVyIiwiaWF0IjoxNjk4Njc1OTY4LCJVc2VyIjp7IklkIjoxLCJVc2VybmFtZSI6ImNvdHRvbiJ9fQ.wPJ_kNrmMpzH_10CHxpDsHkYDyEnIaOul_ABXDiDfuw")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%+v\n", token)
}
