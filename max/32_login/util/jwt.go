package util

import (
	"github.com/golang-jwt/jwt/v4"
	"loginMod/domain"
	"time"
)

const SIGN_KEY = "klasjdnzxcm,nzxn@sdk%sadj;sa"

func JwtCreate(user domain.User) (string, error) {

	//payload 셋팅
	claims := domain.ClaimUser{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Age:      user.Age,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
	}

	//헤더 및 페이로드 생성
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	//서명
	signedString, err := aToken.SignedString([]byte(SIGN_KEY))

	return signedString, err
}

func JwtRead(tokenValue string) (domain.InfoUser, error) {

	var claims domain.ClaimUser
	var infoUser domain.InfoUser

	_, err := jwt.ParseWithClaims(tokenValue, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGN_KEY), nil
	})
	if err != nil {
		return infoUser, err
	}

	infoUser = domain.InfoUser{
		ID:       claims.ID,
		Username: claims.Username,
		Name:     claims.Name,
		Age:      claims.Age,
		Email:    claims.Email,
	}
	return infoUser, nil
}
