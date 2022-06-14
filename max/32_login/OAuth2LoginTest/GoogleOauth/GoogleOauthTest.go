package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var rd *render.Render
var red *redis.Client
var db *gorm.DB
var OAuthConf *oauth2.Config

const (
	SessionID   = "sessionId"
	CallBackURL = "http://localhost:3000/auth/callback"
	// 인증 후 유저 정보를 가져오기 위한 API
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	// 인증 권한 범위. 여기에서는 프로필 정보 권한만 사용
	ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

func main() {
	rd = render.New(render.Options{
		Directory:  "OAuth2LoginTest/template",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := makeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	//레디스 연결
	client, err := initializeRedisClientAPI()
	if err != nil {
		panic(err)
	}
	red = client
	if _, err := red.Ping(context.TODO()).Result(); err != nil {
		panic(err)
	}

	//Mysql 연결
	db, err = ConnectDB()
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	initializeOauth2()

	//if err := db.AutoMigrate(&loginMod.User{}); err != nil {
	//	fmt.Println("User Err")
	//} else {
	//	fmt.Println("User Suc")
	//}

	log.Println("Session Login Started App")
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

func makeWebHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", renderMainView).Methods("GET")
	router.HandleFunc("/auth", renderAuthView).Methods("GET")
	router.HandleFunc("/authUrl", authUrl).Methods("GET")
	router.HandleFunc("/auth/callback", authenticate).Methods("GET")

	return router
}

func renderMainView(w http.ResponseWriter, req *http.Request) {
	rd.HTML(w, http.StatusOK, "main", "b")
}
func renderAuthView(w http.ResponseWriter, req *http.Request) {
	rd.HTML(w, http.StatusOK, "authGoogle", "aa")
}

func authUrl(w http.ResponseWriter, req *http.Request) {
	state := RandToken()
	newUUID, _ := uuid.NewUUID()

	_, err := red.Set(context.TODO(), newUUID.String(), state, time.Minute*5).Result()
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, false)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  SessionID,
		Value: newUUID.String(),
		Path:  "/",
	})

	rd.JSON(w, http.StatusOK, getLoginUrl(state))
}

func authenticate(w http.ResponseWriter, req *http.Request) {

	/*
		state값 가져오기
	*/
	cookie, err := req.Cookie(SessionID)
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, "쿠키 없음")
		return
	}
	defer deleteCookie(w, SessionID)

	state, err := red.Get(context.TODO(), cookie.Value).Result()
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, "레디스에 없음")
		return
	}
	defer red.Del(context.TODO(), cookie.Value)
	if state != req.FormValue("state") {
		rd.JSON(w, http.StatusUnauthorized, "State값 다름")
		return
	}

	/*
		code값 받기, code 사용하여 인증서버에 엑세스 토큰 요청
	*/
	code := req.FormValue("code")
	token, err := OAuthConf.Exchange(context.TODO(), code)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, "code 사용하여 엑세스 토큰 요청 실패")
		return
	}

	/*
		토큰 이용하여 유저정보 요청
	*/
	client := OAuthConf.Client(context.TODO(), token)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint) //UserInfoAPIEndpoint는 유저정보 API URL을 담고있음
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, "토큰이용하여 user정보 요청 실패")
		return
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, "유저정보 읽기 실패")
		return
	}
	//fmt.Println(string(userInfo))
	var user oauthGoogleUser
	err = json.Unmarshal(userInfo, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
	rd.JSON(w, http.StatusOK, user)
	//w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	//http.Redirect(w, req, "/", http.StatusFound)
}

type Success struct {
	Success bool `json:"success"`
}

func deleteCookie(w http.ResponseWriter, cookieName string) {
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Path:   "/",
		Domain: "",
		MaxAge: -1,
	})
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(10.28.3.180:3307)/max?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})
	return db, err
}

func initializeRedisClientAPI() (*redis.Client, error) {
	options := redis.Options{
		Addr:     "10.28.3.180:6379",
		Password: "", //패스워드 없음
		DB:       0,  //기본DB사용
	}

	client := redis.NewClient(&options)
	_, err := client.Ping(context.TODO()).Result()
	return client, err
}

func initializeOauth2() {
	OAuthConf = &oauth2.Config{
		ClientID:     "147137986646-4gmcq6taf5tt2enm6elqgtkuonhc8v0a.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  CallBackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

func getLoginUrl(state string) string {
	return OAuthConf.AuthCodeURL(state)
}

//랜덤 state값
func RandToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

type oauthGoogleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}
