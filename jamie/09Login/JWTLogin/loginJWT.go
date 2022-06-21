package main

import (
	"encoding/json"
	"fmt"
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
	"time"
)

var (
	JWTcookieID = "jwt-token-cookie-id"
	rd          *render.Render
	db          *gorm.DB
	rpath       = regexp.MustCompile(`/userpage`)
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
	n := negroni.Classic()  //-- 3
	n.UseHandler(mux)

	//mysql 연결
	var err error
	db, err = service.ConnectDB()
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
//4. JWT token 생성
//5. 쿠키 안에 토큰값 넣기
func loginHandler(w http.ResponseWriter, r *http.Request) {

	//-- 1
	var loginUser domain.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(loginUser)
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
	tokenvalue, err := service.CreateToken(findUser)
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, false)
		return
	}
	//-- 5
	http.SetCookie(w, &http.Cookie{
		Name:    JWTcookieID,
		Value:   tokenvalue,
		Path:    "/",
		Domain:  "",
		Expires: time.Now().Add(time.Hour * 24),
	})
	rd.JSON(w, http.StatusOK, true)
}

//로그아웃 기능 핸들러
//1. 쿠키 삭제
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	//-- 1
	http.SetCookie(w, &http.Cookie{
		Name:   JWTcookieID,
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

//로그인 체크 핸들러
//1. 쿠키에 있는 토큰 정보 가져오기
//2. 토큰 검증
//2-1. 오류 발생 시 쿠키 삭제
func loginCheckHandler(w http.ResponseWriter, r *http.Request) {

	//-- 1
	cookie, err := r.Cookie(JWTcookieID)
	if err != nil || cookie.Value == "" {
		rd.JSON(w, http.StatusOK, false)
		return
	}
	//-- 2
	infoUser, err := service.VerifyToken(cookie.Value)
	//-- 2-1
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:   JWTcookieID,
			Path:   "/",
			MaxAge: -1,
		})
		rd.JSON(w, http.StatusOK, false)
		return
	}
	rd.JSON(w, http.StatusOK, infoUser)
}

//회원 페이지
func userPageHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "userpage", nil)
}

//회원 인증 미들웨어
//정규식 : https://velog.io/@hsw0194/%EC%A0%95%EA%B7%9C%ED%91%9C%ED%98%84%EC%8B%9D-in-Go
//미들웨어는 핸들러를 감싸는 구조, 핸들러를 파라미터로 전달
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Middleware running")
		path := r.URL.Path
		log.Println("run : ", path)

		authSuccess := true

		switch {
		case rpath.MatchString(path):
			{
				cookie, err := r.Cookie(JWTcookieID)
				if err != nil {
					authSuccess = false
					break
				}
				fmt.Println(cookie.Value)
				_, err = service.VerifyToken(cookie.Value)
				if err != nil {
					http.SetCookie(w, &http.Cookie{
						Name:   JWTcookieID,
						Path:   "/",
						MaxAge: -1,
					})
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
			w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
			log.Println("authentication failed")
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}
	})
}
