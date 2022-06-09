package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var signKey = []byte("SecretCode")

////클레임 내용 - custom claim
type CustomClaims struct {
	TokenUUID            string `json:"tid"`  //토큰 UUID
	Name                 string `json:"name"` //유저 이름
	UserUUID             string `json:"uid"`  //유저 UUID
	jwt.RegisteredClaims        //표준 토큰 Claims
}

func main() {
	createdToken, err := CreateToken(signKey)
	if err != nil {
		fmt.Println("토큰 생성 실패")
	}
	fmt.Println(createdToken)
	VerifyToken(createdToken, string(signKey))

}

func CreateToken(signKey []byte) (string, error) {
	//토큰 생성
	claims := CustomClaims{
		TokenUUID: uuid.NewString(),
		Name:      "jamie",
		UserUUID:  uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Second)), //토큰 만료시간 5초
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//서명 생성 : SignedString() JWT 토큰을 String으로 반환
	s, err := token.SignedString(signKey)
	return s, err
}

func VerifyToken(myToken string, myKey string) {
	token, err := jwt.ParseWithClaims(myToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if token.Valid {
		fmt.Println("token is valid")
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		fmt.Println("유효기간 만료")
	} else {
		fmt.Println("Err:", err)
	}
}
