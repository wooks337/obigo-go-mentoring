package auth

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
	"time"
)

const (
	GoogleCallbackURL = "http://localhost:3000/auth/google/callback"
	OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	ScopeEmail        = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile      = "https://www.googleapis.com/auth/userinfo.profile"

	KakaoCallbackURL = "http://localhost:3000/auth/kakao/callback"
	OauthKakaoUrlAPI = "https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=${REST_API_KEY}&redirect_uri=${REDIRECT_URI}"
)

//oauth2.Config
//Redirect URL : 구글에서 인증완료 후 정보를 callback 할 주소
//ClientID, ClientSecret : 시스템 환경변수에 설정한 값 불러오기
//Scope : 구글 접근 범위 설정(이메일에 접근)
//Endpoint : ???
var GoogleOauthConfig = oauth2.Config{
	RedirectURL:  GoogleCallbackURL,
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{ScopeEmail, ScopeProfile},
	Endpoint:     google.Endpoint,
}

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

//cookie에 일회용 비밀번호 저장
//쿠키 만료 시간 : 현재로부터 24시간
//16byte 짜리 배열을 랜덤하게 채우고 bytes를 string으로 인코딩 -> 이 값을 state 객체로 저장
//http header에 setcookie 설정
func GenerateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour)

	bytes := make([]byte, 16)
	rand.Read(bytes)
	state := base64.URLEncoding.EncodeToString(bytes)

	cookie := &http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie(w, cookie)
	return state
}
