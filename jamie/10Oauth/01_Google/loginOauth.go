package main

import (
	"10Oauth/auth"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	ctx = context.Background()
)

func main() {
	mux := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(mux)      //negroni로 mux 라우터 감싸주기

	log.Println("Started App")
	http.ListenAndServe(":3000", n) //포트 대기
}

func MakeWebHandler() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/auth/kakao/login", kakaoLoginHandler)
	m.HandleFunc("/auth/kakao/callback", kakaoAuthCallback)

	m.HandleFunc("/auth/google/login", googleLoginHandler)
	m.HandleFunc("/auth/google/callback", googleAuthCallback)
	return m
}

//AuthCodeURL로 유저를 어떤 경로로 보내야 하는지 지정 (구글 로그인 경로로 보내야 함-> googleConfig)
//http.Redirect로 요청, 응답, 주소, 상태코드(보내는 이유)
//
//AuthCodeURL에 state 코드를 넣어서 보내야 함
//state 코드란 : CSRF 공격을 막기 위한 토큰 (URL 변조 해킹), 일회용 키 역할  --> 브라우저 쿠키에 임시 키를 심음
//state 객체에 cookie 생성 function을 넣어 전달
//
//이후 callback에 담겨오는 state 객체 값이랑 아래 쿠키에 저장한 state 값을 비교해 일치하면 인증 성공
func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := auth.GenerateStateOauthCookie(w)
	url := auth.GoogleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func kakaoLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := auth.GenerateStateOauthCookie(w)
	url := auth.KakaoOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//유저 구글 로그인이 끝나면 이 콜백핸들러가 호출됨
//==1. 아까 만든 쿠키 호출
//==2. 호출한 쿠키 값과 state 값이 다른 경우, 잘못된 요청으로 간주하여 "/"로 리다이렉트. 오류 내용은 로그로 남기기
//==3. 콜백 시 구글이 r 에 담아주는 코드를 이용하여 유저 정보를 가져온다
//====3-1. 에러 발생 시, "/" 절대 경로로 리다이렉트
//====3-2. 에러 없으면 유저한테 유저정보 보내기
func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprint(w, string(data)) // -- 3-2
}
func kakaoAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate") // -- 1

	if r.FormValue("state") != oauthstate.Value { // -- 2
		log.Printf("invalid google oauth state cookie : %s state : %s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getKakaoUserInfo(r.FormValue("code")) // -- 3
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} // -- 3-1

	fmt.Fprint(w, string(data)) // -- 3-2
}

//구글에서 유저정보 가져오기
func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := auth.GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(auth.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to get userInfo %s\n", err.Error())
	}
	return ioutil.ReadAll(resp.Body)

}

func getKakaoUserInfo(code string) ([]byte, error) {
	token, err := auth.KakaoOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(auth.OauthKakaoUrlAPI)
	if err != nil {
		return nil, fmt.Errorf("Failed to get userInfo %s\n", err.Error())
	}
	return ioutil.ReadAll(resp.Body)
}
