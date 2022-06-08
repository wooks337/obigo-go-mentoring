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
	"log"
	"net/http"
)

type employee struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var rd *render.Render //render 패키지의 Render 구조체 변수 rd
var cli *redis.Client //redis 패키지의 Client 구조체 변수 cli
var ctx = context.Background()

func main() {

	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(m)

	client, err := initialize()
	if err != nil {
		panic(err)
	}
	cli = client
	if _, err := cli.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	log.Println("Started App")
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}

}

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

	m.HandleFunc("/create", CreateSessionHandler).Methods("GET")
	m.HandleFunc("/read", ReadSessionHandler).Methods("GET")
	m.HandleFunc("/delete", DeleteSessionHandler).Methods("GET")
	return m
}

//CreateSession
func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()    // 요청 request url Query 메소드로 파싱
	name := query.Get("name") //빈 객체 생성
	address := query.Get("address")

	newEmp := employee{
		Name:    name,
		Address: address,
	}
	marshal, _ := json.Marshal(&newEmp) //newEmp json 형태로 변환
	newUUID, err := uuid.NewUUID()      //uuid 생성

	//세션 생성
	result, err := cli.Set(ctx, newUUID.String(), marshal, 0).Result()
	if err != nil {
		fmt.Println("생성된 세션 없음", err)
		rd.JSON(w, http.StatusNoContent, err)
		return
	}
	fmt.Println(result) //생성 결과 출력

	//생성된 세션 id 쿠키에 담아 응답으로 전달
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionId",
		Value: newUUID.String(),
		Path:  "/",
	})
	// "/read" 로 리다이렉트
	http.Redirect(w, r, "/read", http.StatusCreated)
}

//ReadSession
func ReadSessionHandler(w http.ResponseWriter, r *http.Request) {

	//요청에서 쿠키값 받아오기
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		rd.JSON(w, http.StatusOK, err.Error())
		return
	}

	var emp employee
	//받아온 쿠키 값의 value 꺼내기
	val, err := cli.Get(ctx, cookie.Value).Result()
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}
	//value = uuid  -> byte 배열로 변경
	//c pcq unmarshal([]byte, v(포인터 형태))
	bytes := []byte(val)
	err = json.Unmarshal(bytes, &emp)

	rd.JSON(w, http.StatusOK, emp)
}

//DeleteSession
func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {

	//요청에서 쿠키값 받아오기
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		rd.JSON(w, http.StatusOK, err.Error())
		return
	}
	res, _ := cli.Del(ctx, cookie.Value).Result()

	rd.JSON(w, http.StatusOK, Success{res == 1})

}

//성공 여부 확인용 Success 구조체
type Success struct {
	Success bool `json:"success"`
}
