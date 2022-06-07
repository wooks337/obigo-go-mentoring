package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
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
	val, err := client.Get(context.TODO(), "aaa").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	//exist
	result, _ := client.Exists(context.TODO(), "xxxxxx").Result()
	fmt.Println(result) //없으면 0
	result, _ = client.Exists(context.TODO(), "aaa").Result()
	fmt.Println(result) //있으면 1

	//delete
	i, _ := client.Del(context.TODO(), "abc", "name").Result()
	fmt.Println(i)

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
