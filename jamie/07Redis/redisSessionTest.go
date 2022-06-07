package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
)

type Employee struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var client *redis.Client
var ctx = context.Background()

func main() {

	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(m)

	json, err := json.Marshal(Employee{Name: "Jamie", Address: "Seoul"})
	if err != nil {
		fmt.Println(err)
	}

	//세션 생성
	err = client.Set(ctx, "emp1", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	err = client.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	////세션 값 읽기
	//val, _ := client.Get(ctx, "emp1").Result()
	//fmt.Println(val)
	//
	//val2, _ := client.Get(ctx, "key1").Result()
	//fmt.Println(val2)

	//모든 키 값 읽기
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("key:", iter.Val())
	}
	if err := iter.Err(); err != nil {
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
