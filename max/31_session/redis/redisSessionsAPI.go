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
	"strconv"
	"time"
)

var rd *render.Render
var red *redis.Client

func main() {
	rd = render.New()
	mux := makeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	client, err := initializeRedisClientAPI()
	if err != nil {
		panic(err)
	}
	red = client
	if _, err := red.Ping(context.TODO()).Result(); err != nil {
		panic(err)
	}

	log.Println("Started App")
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

func makeWebHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/create", createSessionHandler).Methods("GET")
	router.HandleFunc("/read", readSessionHandler).Methods("GET")
	router.HandleFunc("/remove", removeSessionHandler).Methods("GET")
	router.HandleFunc("/update", updateSessionHandler).Methods("GET")
	return router
}

func createSessionHandler(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	name := query.Get("name")
	age, _ := strconv.Atoi(query.Get("age"))

	newUser := user{
		Name: name,
		Age:  age,
	}
	marshal, _ := json.Marshal(&newUser)
	newUUID, err := uuid.NewUUID()
	res, err := red.Set(context.TODO(), newUUID.String(), marshal, time.Minute*5).Result()
	if err != nil {
		fmt.Println("aaa ", err)
		rd.JSON(w, http.StatusNoContent, err)
		return
	}
	fmt.Println("생성 : ", res)
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionId",
		Value: newUUID.String(),
		Path:  "/",
	})

	http.Redirect(w, req, "/read", http.StatusCreated)
}

func readSessionHandler(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("sessionId")
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	var u user

	value, err := red.Get(context.TODO(), cookie.Value).Result()
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	bytesData := []byte(value)
	err = json.Unmarshal(bytesData, &u)
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	rd.JSON(w, http.StatusOK, u)
}

func removeSessionHandler(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("sessionId")
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err.Error())
		return
	}

	result, err := red.Del(context.TODO(), cookie.Value).Result()
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
		return
	}

	rd.JSON(w, http.StatusOK, Success{result == 1})
}

func updateSessionHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("sessionId")
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err.Error())
		return
	}

	var u user
	value, err := red.Get(context.TODO(), cookie.Value).Result()
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	bytesData := []byte(value)
	err = json.Unmarshal(bytesData, &u)
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	u.Age++
	u.Name += "u"
	marshal, _ := json.Marshal(&u)
	_, err = red.Set(context.TODO(), cookie.Value, marshal, time.Minute*5).Result()
	if err != nil {
		fmt.Println("aaa ", err)
		rd.JSON(w, http.StatusNoContent, err)
		return
	}

	http.Redirect(w, req, "/read", http.StatusCreated)
}

type Success struct {
	Success bool `json:"success"`
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

type user struct {
	Name string
	Age  int
}
