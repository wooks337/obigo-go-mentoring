package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {

	//JWT 생성
	signKey := []byte("SecretCode")
	////클레임 내용
	//claims := &jwt.RegisteredClaims{
	//	ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
	//	Issuer:    "test",
	//}

	type CustomClaims struct {
		Name string `json:"name"`
		jwt.RegisteredClaims
	}
	claims := CustomClaims{
		"jamie",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Second)),
			Issuer:    "test",
		},
	}
	//페이로드 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//서명 생성
	s, _ := token.SignedString(signKey)
	//결과
	fmt.Println(s)

	//JWT 검증
	//클레임 파싱
	token2, err := jwt.ParseWithClaims(s, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	if token2.Valid {
		fmt.Println("token is valid")
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		fmt.Println("유효기간 만료")
	} else {
		fmt.Println("Err:", err)
	}
	fmt.Println(token2)
}
