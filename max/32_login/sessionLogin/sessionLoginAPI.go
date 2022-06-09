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
	"regexp"
	"strconv"
)

var rd *render.Render
var red *redis.Client
var db *gorm.DB

const SessionID = "sessionId"

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
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/login-check", loginCheckHandler).Methods("POST")
	router.HandleFunc("/signup", signupPageHandler).Methods("GET")
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("GET")
	router.HandleFunc("/auth", authPageHandler).Methods("GET")
	router.HandleFunc("/auth/profile", myInfoPageHandler).Methods("GET")

	router.Use(authMiddleware)

	return router
}

//var rNum = regexp.MustCompile(`\d`)  // Has digit(s)
var rAuth = regexp.MustCompile(`/auth`) // Contains "abc"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		path := req.URL.Path
		log.Println("미들웨어 작동 :", path) //작업

		authSuccess := true

		switch {
		case rAuth.MatchString(path):
			{
				fmt.Println("auth필요")
				cookie, err := req.Cookie(SessionID)
				if err != nil {
					authSuccess = false
					break
				}
				_, err = util.RedisGet(red, cookie.Value)
				if err != nil {
					authSuccess = false
					break
				}
			}
		default:
			fmt.Println("auth불필요")
		}

		if authSuccess {
			fmt.Println("인증성공 or 인증필요 없음")
			next.ServeHTTP(w, req) // 다음 핸들러 호출
		} else {
			fmt.Println("인증실패")
			http.Redirect(w, req, "/login", http.StatusMovedPermanently)
		}
	})
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

	duplicateCheck := service.UsernameDuplicateCheck(db, signupUser.Username)
	if duplicateCheck == false {
		rd.JSON(w, http.StatusOK, "아이디 중복")
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

func loginHandler(w http.ResponseWriter, req *http.Request) {

	var loginUser loginMod.LoginUser
	err := json.NewDecoder(req.Body).Decode(&loginUser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(loginUser)

	findUser, err := service.FindUserByUsername(db, loginUser.Username)
	if err != nil {
		rd.JSON(w, http.StatusOK, "잘못된 ID")
		return
	}

	passwordCheck := util.PasswordCompare(loginUser.Password, findUser.Password)
	if passwordCheck == false {
		rd.JSON(w, http.StatusOK, "잘못된 PW")
		return
	}

	sessionValue, err := util.RedisSave(red, findUser)
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   SessionID,
		Value:  sessionValue,
		Path:   "/",
		Domain: "",
		MaxAge: 0,
	})
	rd.JSON(w, http.StatusOK, true)
}

func loginCheckHandler(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie(SessionID)
	if err != nil {
		rd.JSON(w, http.StatusOK, false)
		return
	}

	findUser, err := util.RedisGet(red, cookie.Value)
	if err != nil {
		rd.JSON(w, http.StatusOK, false)
		//해당 세션이 없으므로 쿠키 삭제
		http.SetCookie(w, &http.Cookie{
			Name:   SessionID,
			Path:   "/",
			Domain: "",
			MaxAge: -1,
		})
		return
	}

	infoUser := loginMod.InfoUser{
		ID:       findUser.ID,
		Username: findUser.Username,
		Name:     findUser.Name,
		Age:      findUser.Age,
		Email:    findUser.Email,
	}

	rd.JSON(w, http.StatusOK, infoUser)
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {

	//w.Header().Set("Cache-Control", "no-cache, private, max-age=0")

	cookie, err := req.Cookie(SessionID)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
		return
	}

	err = util.RedisDelete(red, cookie.Value)
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, false)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   SessionID,
		Path:   "/",
		Domain: "",
		MaxAge: -1,
	})

	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func authPageHandler(w http.ResponseWriter, req *http.Request) {

	rd.HTML(w, http.StatusOK, "auth", "")
}

func myInfoPageHandler(w http.ResponseWriter, req *http.Request) {

	rd.HTML(w, http.StatusOK, "myInfo", "")
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
