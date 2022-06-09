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
	"log"
	"loginMod"
	"loginMod/service"
	"loginMod/util"
	"net/http"
	"strconv"
)

var rd *render.Render
var red *redis.Client
var db *gorm.DB

func main() {
	rd = render.New(render.Options{
		Directory:  "sessionLogin/template",
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

	//if err := db.AutoMigrate(&loginMod.User{}); err != nil {
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

func makeWebHandler() http.Handler {
	router := mux.NewRouter()
	//router.Handle("/", http.FileServer(http.Dir("sessionLogin/template")))
	router.HandleFunc("/", mainHandler).Methods("GET")
	router.HandleFunc("/login", loginPageHandler).Methods("GET")
	router.HandleFunc("/signup", signupPageHandler).Methods("GET")
	router.HandleFunc("/signup", signupHandler).Methods("POST")

	return router
}

func mainHandler(w http.ResponseWriter, req *http.Request) {

	rd.HTML(w, http.StatusOK, "main", "b")
}
func loginPageHandler(w http.ResponseWriter, req *http.Request) {

	rd.HTML(w, http.StatusOK, "login", "b")
}
func signupPageHandler(w http.ResponseWriter, req *http.Request) {

	rd.HTML(w, http.StatusOK, "join", "b")
}

func signupHandler(w http.ResponseWriter, req *http.Request) {

	var signupUser loginMod.SignupUser
	err := json.NewDecoder(req.Body).Decode(&signupUser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(signupUser)

	age, _ := strconv.Atoi(signupUser.Age)
	passwordHash, err := util.PasswordHash(signupUser.Password)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user := loginMod.User{
		Username: signupUser.Username,
		Password: passwordHash,
		Name:     signupUser.Name,
		Age:      age,
		Email:    signupUser.Email,
	}

	err = service.Signup(db, user)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	rd.JSON(w, http.StatusOK, true)
}

type Success struct {
	Success bool `json:"success"`
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
