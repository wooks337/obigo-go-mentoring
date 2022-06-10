package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gorm.io/gorm"
	"jamie/domain"
	"jamie/service"
	"log"
	"net/http"
)

var (
	rd  *render.Render
	cli *redis.Client
	db  *gorm.DB
	ctx = context.Background()
)

func main() {
	//render 패키지는 기본적으로 확장자를 tmpl로 읽는다
	//html로 된 파일을 읽고 싶으면 옵션을 넣어줘야 한다
	//html과 tmpl 확장자를 둘다 읽도록 옵션 설정

	//render 패키지는 기본적으로 templates에서 찾는다.
	//폴더명을 변경하고 싶을 때 옵션 설정
	rd = render.New(render.Options{
		Directory:  "09Login/templates",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(mux)

	//redis 연결
	client, err := initialize()
	if err != nil {
		panic(err)
	}
	cli = client
	if _, err := cli.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	//mysql 연결
	db, err = service.ConnectDB()
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		log.Println(err)
	}
	//테이블 생성
	//if err := db.AutoMigrate(&domain.User{}); err != nil {
	//	fmt.Println("User Err")
	//} else {
	//	fmt.Println("User Suc")
	//}

	log.Println("Started App")
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

//redis 연결 함수
func initialize() (*redis.Client, error) {
	options := redis.Options{
		Addr:     "10.28.3.180:6379",
		Password: "",
		DB:       0,
	}
	//연결 확인
	client := redis.NewClient(&options)
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to PING Redis: %v", err)
	}
	return client, err
}

//핸들러 등록
func MakeWebHandler() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/", mainHandler).Methods("GET")
	//m.HandleFunc("/idcheck", IDCheckHandler).Methods("GET")
	m.HandleFunc("/signup", signupPageHandler).Methods("GET")
	m.HandleFunc("/signup", signupHandler).Methods("POST")
	m.HandleFunc("/login", loginPageHandler).Methods("GET")
	m.HandleFunc("/login", loginHandler).Methods("POST")
	return m
}

////메인 페이지
func mainHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "index", nil)
}

//회원가입 페이지
func signupPageHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "signup", nil)
}

//func IDCheckHandler(w http.ResponseWriter, r *http.Request) {
//
//	var user domain.User
//	err := json.NewDecoder(r.Body).Decode(&user) //json 형태로 파싱하기 위해 NewDecoder 함수로 요청의 body값을 decode함
//	if err != nil {
//		rd.JSON(w, http.StatusBadRequest, err.Error())
//		return
//	}
//	//아이디 중복 체크
//	idCheck := service.IDCheck(db, user.UserID)
//	if idCheck == false {
//		rd.JSON(w, http.StatusOK, "아이디 중복")
//		return
//	}
//}

//회원가입 핸들러
func signupHandler(w http.ResponseWriter, r *http.Request) {

	//User 구조체 형태의 json을 객체로 받아옴
	var joinuser domain.JoinUser
	//NewDecoder() : 요청 body값으로 들어온 json 데이터를 User구조체 형태로 변경(디코딩)
	err := json.NewDecoder(r.Body).Decode(&joinuser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(joinuser) //데이터 잘 받아왔는지 확인

	//아이디 중복 체크
	idCheck := service.IDCheck(db, joinuser.UserID)
	if idCheck == false {
		rd.JSON(w, http.StatusOK, "아이디 중복")
		return
	}

	//비밀번호 암호화
	pwHash, err := service.PasswordHash(joinuser.Password)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	//변경 데이터(암호화 등) 저장용 데이터 객체화
	user := domain.User{
		UserID:   joinuser.UserID,
		Password: pwHash,
		Name:     joinuser.Name,
		Email:    joinuser.Email,
	}
	//DB에 데이터 저장
	err = service.SignUp(db, user)

	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	rd.JSON(w, http.StatusOK, true)
}

//로그인 페이지
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "login", nil)
}

//로그인 핸들러
func loginHandler(w http.ResponseWriter, r *http.Request) {

	//LoginUser 구조체 형태의 json을 객체로 받아옴
	var loginuser domain.LoginUser

	//1.---사용자가 로그인 화면에서 데이터 입력시 해당 json 데이터를 받아 decode함
	//NewDecoder() : 요청 body값으로 들어온 json 데이터를 LoginUser구조체 형태로 변경(디코딩)
	err := json.NewDecoder(r.Body).Decode(&loginuser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error()) //에러 발생 시, 400오류 반환
		return
	}
	//디코딩한 유저정보 콘솔에서 확인
	fmt.Println(loginuser)

	//2.---DB의 회원정보와 입력받은 로그인 정보를 비교
	findUser, err := service.FindUserByUserid(db, loginuser.UserID)
	if err != nil {
		rd.JSON(w, http.StatusOK, "잘못된 ID")
		return
	}
}
