package service

import (
	"github.com/golang-jwt/jwt/v4"
	"jamie/domain"
	"time"
)

type jwtMethod interface {
	CreateRefreshToken()
	CreateAccessToken()
	VerifyToken()
	CreateReissuanceToken()
}

var signKey = []byte("SecretCode")

//Access 토큰 생성
func CreateAccessToken(user domain.User) (string, error) {

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
	if err != nil {
		return "", err
	}
	return s, nil
}

//Refresh 토큰 생성
func CreateRefreshToken(user domain.User) (string, error) {
	Rclaims := domain.CustomClaims{
		ID:     user.ID,
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Rclaims)
	rt, err := refreshToken.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return rt, nil
}

//토큰 검증
func VerifyToken(myToken string) (domain.InfoUser, error) {

	var claims domain.CustomClaims
	var infoUser domain.InfoUser

	_, err := jwt.ParseWithClaims(myToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return infoUser, err
	}

	infoUser = domain.InfoUser{
		ID:     claims.ID,
		UserID: claims.UserID,
		Name:   claims.Name,
		Email:  claims.Email,
	}
	return infoUser, nil
}
