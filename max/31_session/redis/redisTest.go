package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

func main() {

	client, err := initializeRedisClient()

	if err != nil {
		panic(err)
	}

	//set
	if err := client.Set(context.TODO(), "aaa", "2", -1).Err(); err != nil {
		panic(err)
	}

	//get
	fmt.Println("==Get==")
	val, err := client.Get(context.TODO(), "aaa").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("해당 키가 없음")
	case err != nil:
		fmt.Println("레디스 에러")
	case val == "":
		fmt.Println("value empty")
	default:
		fmt.Println("value is ", val)
	}

	//exist
	fmt.Println("==exists==")
	result, _ := client.Exists(context.TODO(), "xxxxxx").Result()
	fmt.Println(result) //없으면 0
	result, _ = client.Exists(context.TODO(), "aaa").Result()
	fmt.Println(result) //있으면 1

	//delete
	fmt.Println("==Del==")
	i, _ := client.Del(context.TODO(), "abc", "name").Result()
	fmt.Println(i)

	//All search
	fmt.Println("==Scan==")
	keys, cursor, _ := client.Scan(context.TODO(), 0, "*", 0).Result()
	fmt.Println(keys)
	fmt.Println(cursor)

	//update expire
	fmt.Println("==update expire==")
	b, _ := client.Expire(context.TODO(), "aaa", time.Minute*60).Result()
	fmt.Println(b)

	//남은 만료시간 확인
	remainingSeconds, err := client.Do(context.TODO(), "ttl", "sim").Result()
	if err != nil {
		fmt.Println("에러발생", err)
		return
	}
	fmt.Println(remainingSeconds)
}

func initializeRedisClient() (*redis.Client, error) {
	options := redis.Options{
		Addr:     "10.28.3.180:6379",
		Password: "", //패스워드 없음
		DB:       0,  //기본DB사용
	}

	client := redis.NewClient(&options)
	_, err := client.Ping(context.TODO()).Result()
	return client, err
}
