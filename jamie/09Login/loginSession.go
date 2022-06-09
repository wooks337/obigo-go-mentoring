package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jamie/09Login/domain"
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
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(mux)

	//redis
	client, err := initialize()
	if err != nil {
		panic(err)
	}
	cli = client
	if _, err := cli.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	//mysql
	db, err = ConnectDB()
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		log.Println(err)
	}

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

	//m.Handle("/", http.FileServer(http.Dir("templates")))		//왜 안됨?
	m.HandleFunc("/", mainHandler).Methods("GET")
	m.HandleFunc("/login", loginPageHandler).Methods("GET")
	m.HandleFunc("/signup", signupPageHandler).Methods("GET")
	m.HandleFunc("/signup", signupHandler).Methods("POST")
	return m
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "index", nil)
}
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "login", "")
}
func signupPageHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "signup", "")
}

//회원가입 핸들러
func signupHandler(w http.ResponseWriter, r *http.Request) {

	var user domain.User                         //user struct 형태의 json을 객체로 받아서
	err := json.NewDecoder(r.Body).Decode(&user) //json 형태로 파싱하기 위해 NewDecoder 함수로 요청의 body값을 decode함
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	err = SignUp(db, user)

	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	rd.JSON(w, http.StatusOK, true)

}

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
