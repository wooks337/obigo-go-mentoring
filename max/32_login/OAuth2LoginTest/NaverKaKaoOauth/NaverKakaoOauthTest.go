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
var KakaoOAuthConf *oauth2.Config
var NaverOAuthConf *oauth2.Config

const (
	SessionID        = "sessionId"
	KakaoCallBackURL = "http://localhost:3000/auth/kakao/callback"
	NaverCallBackURL = "http://localhost:3000/auth/naver/callback"
	// 인증 후 유저 정보를 가져오기 위한 API
	UserInfoKakaoAPIEndpoint = "https://kapi.kakao.com/v2/user/me"
	UserInfoNaverAPIEndpoint = "https://openapi.naver.com/v1/nid/me"
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
	router.HandleFunc("/auth/kakao/callback", kakaoAuthenticate).Methods("GET")
	router.HandleFunc("/auth/naver/callback", naverAuthenticate).Methods("GET")

	return router
}

func renderMainView(w http.ResponseWriter, req *http.Request) {
	rd.HTML(w, http.StatusOK, "main", "b")
}
func renderAuthView(w http.ResponseWriter, req *http.Request) {
	rd.HTML(w, http.StatusOK, "authNaverKakao", "aa")
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

	kakaoAuthUrl := getLoginUrl(KakaoOAuthConf, state)
	naverAuthUrl := getLoginUrl(NaverOAuthConf, state)

	res := make(map[string]string)
	res["kakao"] = kakaoAuthUrl
	res["naver"] = naverAuthUrl

	rd.JSON(w, http.StatusOK, res)
}

func kakaoAuthenticate(w http.ResponseWriter, req *http.Request) {

	//state값 가져오기
	err := validationToState(w, req)
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err)
		return
	}

	//code값 받기, code 사용하여 인증서버에 엑세스 토큰 요청
	userInfo, err := getOAuthUserInfoFromCode(w, req, KakaoOAuthConf)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err)
		return
	}

	var user oauthKakaoUser
	err = json.Unmarshal(userInfo, &user)
	if err != nil {
		rd.JSON(w, http.StatusOK, err.Error())
		return
	}

	rd.JSON(w, http.StatusOK, user)
}

func naverAuthenticate(w http.ResponseWriter, req *http.Request) {

	//state값 가져오기
	err := validationToState(w, req)
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err)
		return
	}

	//code값 받기, code 사용하여 인증서버에 엑세스 토큰 요청
	userInfo, err := getOAuthUserInfoFromCode(w, req, NaverOAuthConf)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err)
		return
	}

	var user oauthNaverUser
	err = json.Unmarshal(userInfo, &user)
	if err != nil {
		rd.JSON(w, http.StatusOK, err.Error())
		return
	}

	rd.JSON(w, http.StatusOK, user)
}

func getOAuthUserInfoFromCode(w http.ResponseWriter, req *http.Request, oauthConfig *oauth2.Config) ([]byte, error) {

	//code값 이용하여 토큰발급
	code := req.FormValue("code")
	token, err := oauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, fmt.Errorf("code 사용하여 엑세스 토큰 요청 실패")
	}

	//토큰 이용하여 유저정보 요청
	var userInfoResp *http.Response
	client := oauthConfig.Client(context.TODO(), token)
	switch oauthConfig {
	case KakaoOAuthConf:
		userInfoResp, err = client.Get(UserInfoKakaoAPIEndpoint)
	case NaverOAuthConf:
		userInfoResp, err = client.Get(UserInfoNaverAPIEndpoint)
	}
	if err != nil {
		return nil, fmt.Errorf("토큰이용하여 user정보 요청 실패")
	}

	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return nil, fmt.Errorf("유저정보 읽기 실패")
	}
	return userInfo, nil
}

func validationToState(w http.ResponseWriter, req *http.Request) error {

	cookie, err := req.Cookie(SessionID)
	if err != nil {
		return fmt.Errorf("쿠키없음")
	}
	defer deleteCookie(w, SessionID)

	state, err := red.Get(context.TODO(), cookie.Value).Result()
	if err != nil {
		return fmt.Errorf("레디스에 없음")

	}
	defer red.Del(context.TODO(), cookie.Value)
	if state != req.FormValue("state") {
		return fmt.Errorf("State값 다름")
	}
	return nil
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
	KakaoOAuthConf = &oauth2.Config{
		ClientID:     "d62586d693310611ca4dac9525af96e7",
		ClientSecret: "DCpVOpE2B2DP32M9XiPUaV3kxEHz9A9Y",
		RedirectURL:  KakaoCallBackURL,
		Scopes:       []string{"profile_nickname", "profile_image"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
	}

	NaverOAuthConf = &oauth2.Config{
		ClientID:     "peb88WhrIfRq7k4EPow4",
		ClientSecret: "LYwZBAejND",
		RedirectURL:  NaverCallBackURL,
		Scopes:       []string{"name", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://nid.naver.com/oauth2.0/authorize",
			TokenURL: "https://nid.naver.com/oauth2.0/token",
		},
	}
}

func getLoginUrl(config *oauth2.Config, state string) string {
	return config.AuthCodeURL(state)
}

//랜덤 state값
func RandToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

type oauthKakaoUser struct {
	Id          int               `json:"id"`
	ConnectedAt time.Time         `json:"connected_at"`
	Properties  map[string]string `json:"properties"`
}

/*
{
    "id":2285551298,
    "connected_at":"2022-06-14T02:33:18Z",
    "properties":{
        "nickname":"심종현",
        "profile_image":"http://k.kakaocdn.net/dn/J9LSK/btrANSZ5APH/t0YdKXsdAnSVK9IWP31XGk/img_640x640.jpg",
        "thumbnail_image":"http://k.kakaocdn.net/dn/J9LSK/btrANSZ5APH/t0YdKXsdAnSVK9IWP31XGk/img_110x110.jpg"
        },
    "kakao_account":{
        "profile_nickname_needs_agreement":false,
        "profile_image_needs_agreement":false,
        "profile":{
            "nickname":"심종현",
            "thumbnail_image_url":"http://k.kakaocdn.net/dn/J9LSK/btrANSZ5APH/t0YdKXsdAnSVK9IWP31XGk/img_110x110.jpg",
            "profile_image_url":"http://k.kakaocdn.net/dn/J9LSK/btrANSZ5APH/t0YdKXsdAnSVK9IWP31XGk/img_640x640.jpg",
            "is_default_image":false
            }
    }
}
*/

type oauthNaverUser struct {
	ResultCode string            `json:"resultcode"`
	Message    string            `json:"message"`
	Response   map[string]string `json:"response"`
}

/*
{
	resultcode: "00",
	message: "success",
	response: {
		email: "whdgus5289@naver.com",
		id: "urfZFDXcFbteTWWoGrb3JJa7imLQAhRxAhaum57ydH0",
		name: "심종현",
	},
}
*/
