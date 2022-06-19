package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"jamie/domain"
	"jamie/service"
	"log"
	"net/http"
	"regexp"
)

var (
	rd    *render.Render
	cli   *redis.Client
	db    *gorm.DB
	ctx   = context.Background()
	rpath = regexp.MustCompile(`/userpage`)
)

//main 함수
//1. render 변수 선언 : html 확장자 옵션 처리
//2. 사용자 핸들러 함수 담기
//3. negroni 기본 핸들러 선언 + 사용자 핸들러 담기
func main() {
	rd = render.New(render.Options{ //-- 1
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := MakeWebHandler() //-- 2

	n := negroni.Classic() //-- 3
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

//사용자 핸들러 함수
//return 값 : 핸들러 인스턴스
func MakeWebHandler() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/", mainHandler).Methods("GET")
	m.HandleFunc("/signup", signupPageHandler).Methods("GET")
	m.HandleFunc("/signup", signupHandler).Methods("POST")
	m.HandleFunc("/login", loginPageHandler).Methods("GET")
	m.HandleFunc("/login", loginHandler).Methods("POST")
	m.HandleFunc("/logincheck", loginCheckHandler).Methods("POST")
	m.HandleFunc("/logout", logoutHandler).Methods("GET")
	m.HandleFunc("/userpage", userPageHandler).Methods("GET")

	m.Use(authMiddleware)
	//m.Use(DummyMiddleware)
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

//회원가입 핸들러
//1. 회원가입 창에서 입력한 데이터 받아오기
//2. 아이디 중복 체크
//3. 비밀번호 암호화
//4. DB에 데이터 저장
func signupHandler(w http.ResponseWriter, r *http.Request) {

	//-- 1
	var joinuser domain.JoinUser
	err := json.NewDecoder(r.Body).Decode(&joinuser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(joinuser) //데이터 잘 받아왔는지 확인

	//-- 2
	idCheck := service.IDCheck(db, joinuser.UserID)
	if idCheck == false {
		rd.JSON(w, http.StatusOK, "아이디 중복")
		return
	}

	//-- 3
	pw, err := bcrypt.GenerateFromPassword([]byte(joinuser.Password), bcrypt.DefaultCost)
	pwHash := string(pw)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	//-- 4
	user := domain.User{
		UserID:   joinuser.UserID,
		Password: pwHash,
		Name:     joinuser.Name,
		Email:    joinuser.Email,
	}
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
//1. 회원가입 창에서 입력한 데이터 받아오기
//2. DB의 아이디와 입력받은 아이디 정보를 비교
//3. 로그인 시 입력한 비밀번호와 db저장 비밀번호 비교
//4. 사용자 세션 생성
//5. 사용자 세션 쿠키에 저장
func loginHandler(w http.ResponseWriter, r *http.Request) {

	//-- 1
	var loginUser domain.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	//-- 2
	findUser, err := service.FindUserByUserid(db, loginUser.UserID)
	if err != nil {
		rd.JSON(w, http.StatusOK, "wrong ID")
		return
	}
	log.Println("==1.raw", loginUser.Password)
	log.Println("==2.hash", findUser.Password)

	//-- 3
	checkPassword := service.CheckHashPassword(findUser.Password, loginUser.Password)
	if checkPassword == false {
		rd.JSON(w, http.StatusOK, "wrong PW")
		return
	}

	//-- 4
	session, err := service.RedisSessionCreate(cli, findUser)
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
	}
	//-- 5
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: session,
		Path:  "/",
	})
	rd.JSON(w, http.StatusOK, true)

}

//로그아웃 기능 핸들러
//1. 쿠키에 있는 세션 정보 가져오기
//2. 세션 정보 삭제
func logoutHandler(w http.ResponseWriter, r *http.Request) {

	//-- 1
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	//-- 2
	err = service.RedisSessionDelete(cli, cookie.Value)
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, nil)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: "sessionID",
		Path: "/",
	})
	http.Redirect(w, r, "/", http.StatusOK)
}

//로그인 체크 핸들러
//1. 쿠키에 있는 세션 정보 가져오기
//2. 세션 키 값으로 현재 로그인 한 유저 정보 조회
//2-1. 해당하는 세션이 없을 경우 쿠키 삭제
//3. 조회한 유저정보 구조체에 넣어 반환하기
func loginCheckHandler(w http.ResponseWriter, r *http.Request) {

	//-- 1
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		rd.JSON(w, http.StatusOK, false)
		return
	}
	//-- 2
	findUser, err := service.RedisSessionRead(cli, cookie.Value)
	//-- 2-1
	if err != nil {
		rd.JSON(w, http.StatusOK, false)
		http.SetCookie(w, &http.Cookie{
			Name:   "sessionID",
			Path:   "/",
			Domain: "",
			MaxAge: -1,
		})
		return
	}
	//-- 3
	infoUser := domain.InfoUser{
		ID:     findUser.ID,
		UserID: findUser.UserID,
		Name:   findUser.Name,
		Email:  findUser.Email,
	}
	rd.JSON(w, http.StatusOK, infoUser)
}

//회원 페이지
func userPageHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "userpage", nil)
}

////미들웨어 테스트
//https://www.notion.so/Gorilla-7621ae82b7df423fb6033919612b96db#6ec6dc0c13b74ac3a8809fe91cd4c797
//https://eli.thegreenplace.net/2021/life-of-an-http-request-in-a-go-server/
//func DummyMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("Middleware Test Dummy")
//		next.ServeHTTP(w, r)
//		log.Println("Middle execute")
//	})
//}

//회원 인증 미들웨어
//정규식 : https://velog.io/@hsw0194/%EC%A0%95%EA%B7%9C%ED%91%9C%ED%98%84%EC%8B%9D-in-Go
//미들웨어는 핸들러를 감싸는 구조, 핸들러를 파라미터로 전달
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Middleware running")
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		path := r.URL.Path
		log.Println("run : ", path)

		authSuccess := true

		switch {
		case rpath.MatchString(path):
			{
				cookie, err := r.Cookie("sessionID")

				if err != nil {
					authSuccess = false
					break
				}
				_, err = service.RedisSessionRead(cli, cookie.Value)
				if err != nil {
					authSuccess = false
					break
				}
			}
		default:
			log.Println("authentication no needed")
		}
		if authSuccess {
			log.Println("authentication success")
			next.ServeHTTP(w, r)
		} else {
			log.Println("authentication failed")
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}
	})
}
