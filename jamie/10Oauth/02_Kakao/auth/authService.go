package auth

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
)

const (
	KakaoCallbackURL = "http://localhost:3000/auth/kakao/callback"
	OauthKakaoUrlAPI = "https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=${REST_API_KEY}&redirect_uri=${REDIRECT_URI}"
)

var KakaoOauthConfig = oauth2.Config{
	RedirectURL:  KakaoCallbackURL,
	ClientID:     "-",
	ClientSecret: "-",
	Scopes:       []string{"profile_nickname", "account_email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://kauth.kakao.com/oauth/authorize",
		TokenURL: "https://kauth.kakao.com/oauth/token",
	},
}

func RandToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)

}
