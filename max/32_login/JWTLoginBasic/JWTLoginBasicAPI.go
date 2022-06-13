package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"loginMod/domain"
	"loginMod/service"
	"loginMod/util"
	"net/http"
	"regexp"
	"strconv"
)

var rd *render.Render

var db *gorm.DB

const JWT_COOKIE_ID = "jwt-token"

func main() {
	rd = render.New(render.Options{
		Directory:  "JWTLoginBasic/static/template",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := makeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	//Mysql 연결
	var err error
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

	log.Println("JWT Login Basic Started App")
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

func makeWebHandler() http.Handler {
	router := mux.NewRouter()
	router.Handle("/static/{rest}/{file}", http.StripPrefix("/static/", http.FileServer(http.Dir("JWTLoginBasic/static/"))))
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

		path := req.URL.Path
		log.Println("미들웨어 작동 :", path) //작업

		authSuccess := true

		switch {
		case rAuth.MatchString(path):
			{
				fmt.Println("auth필요")
				cookie, err := req.Cookie(JWT_COOKIE_ID)
				if err != nil {
					authSuccess = false
					break
				}
				fmt.Println(cookie.Value)
				_, err = util.JwtRead(cookie.Value)
				if err != nil {
					http.SetCookie(w, &http.Cookie{
						Name:   JWT_COOKIE_ID,
						Path:   "/",
						Domain: "",
						MaxAge: -1,
					})
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
			w.Header().Set("Cache-Control", "no-cache, private, max-age=0")

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

	var signupUser domain.SignupUser
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

	user := domain.User{
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

	var loginUser domain.LoginUser
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

	jwtValue, err := util.JwtCreate(findUser)
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, false)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   JWT_COOKIE_ID,
		Value:  jwtValue,
		Path:   "/",
		Domain: "",
		MaxAge: 60 * 60, //1시간
	})
	rd.JSON(w, http.StatusOK, true)
}

func loginCheckHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")

	cookie, err := req.Cookie(JWT_COOKIE_ID)
	if err != nil || cookie.Value == "" {
		rd.JSON(w, http.StatusOK, false)
		return
	}

	infoUser, err := util.JwtRead(cookie.Value)
	if err != nil {

		http.SetCookie(w, &http.Cookie{
			Name:   JWT_COOKIE_ID,
			Path:   "/",
			Domain: "",
			MaxAge: -1,
		})
		if errors.Is(err, jwt.ErrTokenExpired) {
			//	rd.JSON(w, http.StatusOK, false) //토큰만료
			fmt.Println("JWT 기간 만료")
			return
		}
		rd.JSON(w, http.StatusOK, false)
		return
	}

	rd.JSON(w, http.StatusOK, infoUser)
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie(JWT_COOKIE_ID)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   cookie.Name,
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
