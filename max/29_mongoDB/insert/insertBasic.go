package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"mongoTest/connect"
)

func main() {

	mongoDB, err := connect.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer func() {
		mongoDB.Disconnect(context.TODO())
	}()

	bookCollection := mongoDB.Database("maxSample").Collection("book")
	res, err := bookCollection.InsertOne(context.TODO(), bson.D{
		{"company", "기아"}, {"content", "content05"}, {"intarray", []int{1, 2, 3}},
		{"likes", 24}, {"title", "Naruto"}, {"writer", "ninza"},
	})
	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println("Suc")
	}
	fmt.Println(res)
}
