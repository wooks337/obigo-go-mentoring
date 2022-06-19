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
	"io/ioutil"
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

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func main() {
	//render 패키지는 기본적으로 확장자를 tmpl로 읽는다
	//html로 된 파일을 읽고 싶으면 옵션을 넣어줘야 한다
	//html과 tmpl 확장자를 둘다 읽도록 옵션 설정
	//
	//render 패키지는 기본적으로 templates에서 찾는다.
	//폴더명을 변경하고 싶을 때 옵션 설정
	rd = render.New(render.Options{
		Directory:  "templates",
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
	m.HandleFunc("/signup", signupPageHandler).Methods("GET")
	m.HandleFunc("/signup", signupHandler).Methods("POST")
	m.HandleFunc("/login", loginPageHandler).Methods("GET")
	m.HandleFunc("/login", loginHandler).Methods("POST")
	m.HandleFunc("/logincheck", loginCheckHandler).Methods("POST")
	m.HandleFunc("/logout", logoutHandler).Methods("GET")
	m.HandleFunc("/userpage", userPageHandler).Methods("GET")

	m.HandleFunc("/auth/google/login", googleLoginHandler)
	m.HandleFunc("/auth/google/callback", googleAuthCallback)

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
	pwHash, err := service.HashPassword(joinuser.Password)
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
	var loginUser domain.LoginUser

	//1.---사용자가 로그인 화면에서 데이터 입력시 해당 json 데이터를 받아 decode함
	//NewDecoder() : 요청 body값으로 들어온 json 데이터를 LoginUser 구조체 형태로 변경(디코딩)
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error()) //에러 발생 시, 400오류 반환
		return
	}
	//디코딩한 로그인창 입력 유저정보 콘솔에서 확인
	log.Println("==1.raw", loginUser.Password)

	//2.---DB의 회원정보와 입력받은 로그인 정보를 비교
	//finduser = db에 저장된 회원정보
	findUser, err := service.FindUserByUserid(db, loginUser.UserID)
	if err != nil {
		rd.JSON(w, http.StatusOK, "wrong ID")
		return
	}
	log.Println("==2.hash", findUser.Password)

	//로그인 시 입력한 비밀번호와 db저장 비밀번호 비교
	checkPassword := service.CheckHashPassword(findUser.Password, loginUser.Password)
	if checkPassword == false {
		rd.JSON(w, http.StatusOK, "wrong PW")
		return
	}

	//3.---로그인 시 사용자 세션 생성
	session, err := service.RedisSessionCreate(cli, findUser)
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
	}
	//쿠키에 세션 정보 저장
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: session,
		Path:  "/",
	})
	rd.JSON(w, http.StatusOK, true) //여기에 json 리턴값 nil로 줘서 결과값 계속 null 나옴... true로 줘야 정상적으로 alert 띄움~

}

//로그아웃 핸들러
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	//쿠키에 있는 세션 정보 가져오기
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	//세션 삭제하기
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
func loginCheckHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		rd.JSON(w, http.StatusOK, false)
		return
	}
	findUser, err := service.RedisSessionRead(cli, cookie.Value)
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

//구글 로그인 핸들러
func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := service.GenerateStateOauthCookie(w)
	url := service.GoogleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	//<<session>>

	//  state := auth.RandToken()
	//	newUUID, _ := uuid.NewUUID()
	//
	//	//Redis Session Create
	//	//=== 1 세션에 state 값(임의값) 저장
	//	//=== 2 생성된 세션 id 쿠키에 담아 응답으로 전달
	//	res, err := cli.Set(ctx, newUUID.String(), state, time.Hour*1).Result() //---1
	//	if err != nil {
	//		rd.JSON(w, http.StatusInternalServerError, err.Error())
	//	}
	//	fmt.Println(res)
	//
	//	http.SetCookie(w, &http.Cookie{ //---2
	//		Name:  "sessionID",
	//		Value: newUUID.String(),
	//		Path:  "/",
	//	})
	//
	//	url := service.GoogleOauthConfig.AuthCodeURL(state)
	//  rd.JSON(w, http.StatusOK, url)

}

//구글 콜백 핸들러
func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	//<<session>>

	//=== 1. 아까 만든 쿠키 호출
	//=== 2. 쿠키에서 state값 꺼내기
	//=== 3. 호출한 쿠키 값과 state 값이 다른 경우, 잘못된 요청으로 간주하여 "/"로 리다이렉트. 오류 내용은 로그로 남기기
	//	oauthstate, _ := r.Cookie("oauthstate") // -- 1
	//	state, _ := cli.Get(context.TODO(), oauthstate.Value).Result() //--2
	//	if r.FormValue("state") != state { // -- 2
	//		log.Printf("invalid google oauth state cookie : %s state : %s\n", oauthstate.Value, r.FormValue("state"))
	//		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//		return
	//	}
	//
	//	data, err := getGoogleUserInfo(r.FormValue("code")) // -- 3
	//	if err != nil {
	//		log.Println(err.Error())
	//		rd.JSON(w, http.StatusBadRequest, "유저정보 읽기 실패")
	//		return
	//	}
	//
	//	fmt.Println(string(data))
	//
	//	//GoogleUser 구조체 형태의 json을 객체로 받아옴
	//	var googleuser domain.GoogleUser
	//
	//	err = json.Unmarshal(data, &googleuser)
	//	if err != nil {
	//		rd.JSON(w, http.StatusBadRequest, err.Error())
	//		return
	//	}
	//
	//	//변경 데이터 저장용 데이터 객체화
	//	user2 := domain.User{
	//		UserID: googleuser.ID,
	//		Name:   googleuser.Name,
	//		Email:  googleuser.Email,
	//	}
	//
	//	//DB에 데이터 저장
	//	err = service.SignUp(db, user2)
	//	if err != nil {
	//		rd.JSON(w, http.StatusBadRequest, err.Error())
	//		return
	//	}
	//	rd.HTML(w, http.StatusOK, "index", nil)
	//}

	oauthstate, _ := r.Cookie("oauthstate") // -- 1

	if r.FormValue("state") != oauthstate.Value { // -- 2
		log.Printf("invalid google oauth state cookie : %s state : %s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getGoogleUserInfo(r.FormValue("code")) // -- 3
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} // -- 3-1

	fmt.Println(string(data))

	//GoogleUser 구조체 형태의 json을 객체로 받아옴
	var googleuser domain.GoogleUser

	err = json.Unmarshal(data, &googleuser)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	//변경 데이터(암호화 등) 저장용 데이터 객체화
	user2 := domain.User{
		UserID: googleuser.ID,
		Name:   googleuser.Name,
		Email:  googleuser.Email,
	}

	//DB에 데이터 저장
	err = service.SignUp(db, user2)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	rd.HTML(w, http.StatusOK, "index", nil)
}

//구글에서 유저정보 가져오기
func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := service.GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to get userInfo %s\n", err.Error())
	}
	return ioutil.ReadAll(resp.Body)

}

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