package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
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

	m.HandleFunc("/auth/google/login", googleLoginHandler)
	m.HandleFunc("/auth/google/callback", googleAuthCallback)

	return m
}

////메인 페이지
func mainHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "index", nil)
}

//구글 로그인 핸들러
func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := service.GenerateStateOauthCookie(w)
	newUUID, _ := uuid.NewUUID()

	result, err := cli.Set(ctx, newUUID.String(), state, 0).Result()
	if err != nil {
		fmt.Println("생성된 세션 없음", err)
		rd.JSON(w, http.StatusNoContent, err)
		return
	}
	fmt.Println(result) //생성 결과 출력

	////생성된 세션 id 쿠키에 담아 응답으로 전달
	//http.SetCookie(w, &http.Cookie{
	//	Name:  "sessionId",
	//	Value: newUUID.String(),
	//	Path:  "/",
	//})

	url := service.GoogleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

//구글 콜백 핸들러
//=== 1. 아까 만든 쿠키 호출
//=== 2. 쿠키에서 state값 꺼내기
//=== 3. 호출한 쿠키 값과 state 값이 다른 경우, 잘못된 요청으로 간주하여 "/"로 리다이렉트. 오류 내용은 로그로 남기기
func googleAuthCallback(w http.ResponseWriter, r *http.Request) {

	// -- 1
	oauthstate, err := r.Cookie("sessionID")
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, "no cookie")
		return
	}
	////-- 2
	//state, err := cli.Get(ctx, oauthstate.Value).Result()
	//if err != nil {
	//	rd.JSON(w, http.StatusUnauthorized, "no state")
	//}
	// -- 3
	if r.FormValue("state") != oauthstate.Value {
		log.Printf("invalid google oauth state cookie : %s state : %s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := service.GetGoogleUserInfo(r.FormValue("code")) // -- 3
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintln(w, string(data))

	//GoogleUser 구조체 형태의 json을 객체로 받아옴
	var googleuser domain.GoogleUser

	err = json.Unmarshal(data, &googleuser)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//변경 데이터 저장용 데이터 객체화
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
