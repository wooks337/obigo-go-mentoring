package main

import (
	"02_Kakao/auth"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	rd  *render.Render
	cli *redis.Client
	ctx = context.Background()
)

func main() {
	rd = render.New(render.Options{
		Directory:  "public",
		Extensions: []string{".html", ".tmpl"},
	})

	//redis 연결
	client, err := initialize()
	if err != nil {
		panic(err)
	}
	cli = client
	if _, err := cli.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	mux := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(mux)      //negroni로 mux 라우터 감싸주기

	log.Println("Started App")
	http.ListenAndServe(":3000", n) //포트 대기
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

func MakeWebHandler() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/auth/kakao/login", kakaoLoginHandler)
	m.HandleFunc("/auth/kakao/callback", kakaoAuthCallback)
	return m
}

func kakaoLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := auth.RandToken()
	newUUID, _ := uuid.NewUUID()

	//Redis Session Create
	//=== 1 세션에 state 값(임의값) 저장
	//=== 2 생성된 세션 id 쿠키에 담아 응답으로 전달
	res, err := cli.Set(ctx, newUUID.String(), state, time.Minute*1).Result() //---1
	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, err.Error())
	}
	fmt.Println(res)

	http.SetCookie(w, &http.Cookie{ //---2
		Name:  "sessionID",
		Value: newUUID.String(),
		Path:  "/",
	})

	url := auth.KakaoOauthConfig.AuthCodeURL(state)
	//	rd.JSON(w, http.StatusOK, url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//=== 1. 아까 만든 쿠키 호출
//=== 2. 쿠키에서 state값 꺼내기
//=== 3. 호출한 쿠키 값과 state 값이 다른 경우, 잘못된 요청으로 간주하여 "/"로 리다이렉트. 오류 내용은 로그로 남기기

func kakaoAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("sessionID")                         // -- 1
	state, _ := cli.Get(context.TODO(), oauthstate.Value).Result() //--2
	if r.FormValue("state") != state {                             // -- 3
		log.Printf("invalid google oauth state cookie : %s state : %s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
}

func getUserInfo(code string) ([]byte, error) {
	token, err := auth.KakaoOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(auth.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to get userInfo %s\n", err.Error())
	}
	return ioutil.ReadAll(resp.Body)

}
