package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.28.3.180:6379",
		Password: "",
		DB:       0,
	})
	//연결 확인
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to PING Redis: %v", err)
	}
	fmt.Println(pong)
	//데이터 생성
	err = client.Set(ctx, "name", "jamie", 0).Err()
	if err != nil {
		panic(err)
	}
	//데이터 읽기
	val, _ := client.Get(ctx, "name").Result()
	fmt.Println(val)
	//데이터 삭제
	//del, _ := client.Del(ctx, "name").Result()
	//fmt.Println(del)

	iter := client.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
}
