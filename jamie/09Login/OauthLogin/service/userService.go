package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jamie/domain"
	"net/http"
	"os"
	"time"
)

const (
	CallbackURL       = "http://localhost:3000/auth/google/callback"
	OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	ScopeEmail        = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile      = "https://www.googleapis.com/auth/userinfo.profile"
)

//mysql 서버 연결 함수
func ConnectDB() (*gorm.DB, error) {
	//dsn := "root:jamiekim@(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
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

//아이디 중복 체크 함수
func IDCheck(db *gorm.DB, userid string) bool {
	findID := domain.User{}
	res := db.Model(&domain.User{}).First(&findID, "userid = ?", userid)
	if res.Error != nil {
		return true
	} else {
		return false
	}
}

//비밀번호 암호화 함수
//https://bourbonkk.tistory.com/64
//https://jeong-dev-blog.tistory.com/2
//pwHash, _ := bcrypt.GenerateFromPassword([]byte(), bcrypt.DefaultCost)
//[]byte 자료형의 해시 반환 -> 해시 반환값을 string 변환 후 DB 저장
func HashPassword(password string) (string, error) {
	pw := []byte(password)

	pwHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", pwHash), nil
}

//userid로 회원정보 찾기 함수
func FindUserByUserid(db *gorm.DB, userid string) (domain.User, error) {
	//User 구조체에 회원조회 정보 담아서 에러랑 같이 반환하기
	findUser := domain.User{}
	res := db.Model(&domain.User{}).First(&findUser, "user_id = ?", userid)
	return findUser, res.Error
}

//로그인 시 비밀번호 일치 확인 함수
func CheckHashPassword(hashVal, userPw string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}

//oauth2.Config
//Redirect URL : 구글에서 인증완료 후 정보를 callback 할 주소
//ClientID, ClientSecret : 시스템 환경변수에 설정한 값 불러오기
//Scope : 구글 접근 범위 설정(이메일에 접근)
//Endpoint : ???
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
