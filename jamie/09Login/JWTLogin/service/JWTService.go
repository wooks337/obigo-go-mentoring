package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"jamie/domain"
	"time"
)

var signKey = []byte("SecretCode")

//토큰 생성
func CreateToken(user domain.User) (string, error) {

	claims := domain.CustomClaims{
		ID:     user.ID,
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Second)), //토큰 만료시간 5초
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	//서명 생성 : SignedString() JWT 토큰을 String으로 반환
	s, err := token.SignedString(signKey)
	return s, err
}

//토큰 검증
func VerifyToken(myToken string) (domain.InfoUser, error) {

	var claims domain.CustomClaims
	var infoUser domain.InfoUser

	token, err := jwt.ParseWithClaims(myToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return infoUser, err
	}

	if token.Valid {
		fmt.Println("token is valid")
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		fmt.Println("유효기간 만료")
	} else {
		fmt.Println("Err:", err)
	}

	infoUser = domain.InfoUser{
		ID:     claims.ID,
		UserID: claims.UserID,
		Name:   claims.Name,
		Email:  claims.Email,
	}
	return infoUser, nil
}
