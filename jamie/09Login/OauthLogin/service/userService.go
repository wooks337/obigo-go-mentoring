package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"jamie/domain"
	"os"
)

const (
	CallbackURL       = "http://localhost:3000/auth/google/callback"
	OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	ScopeEmail        = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile      = "https://www.googleapis.com/auth/userinfo.profile"
)

//mysql 서버 연결 함수
func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db, err
}

//db 데이터 저장 함수
func SignUp(db *gorm.DB, user domain.User) error {
	res := db.Create(&user)
	return res.Error
}

//oauth2.Config
//Redirect URL : 구글에서 인증완료 후 정보를 callback 할 주소
//ClientID, ClientSecret : 시스템 환경변수에 설정한 값 불러오기
//Scope : 구글 접근 범위 설정(이메일에 접근)

var GoogleOauthConfig = oauth2.Config{
	RedirectURL:  CallbackURL,
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{ScopeEmail, ScopeProfile},
	Endpoint:     google.Endpoint,
}

//cookie에 일회용 비밀번호 저장
//쿠키 만료 시간 : 현재로부터 24시간
//16byte 짜리 배열을 랜덤하게 채우고 bytes를 string으로 인코딩 -> 이 값을 state 객체로 저장
//http header에 setcookie 설정
func GenerateRandomToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

//구글에서 유저정보 가져오기
func GetGoogleUserInfo(code string) ([]byte, error) {
	token, err := GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}
	client := GoogleOauthConfig.Client(ctx, token)
	resp, err := client.Get(OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to get userInfo %s\n", err.Error())
	}
	return ioutil.ReadAll(resp.Body)
}
