package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	////Index 조회
	//list, err := bookCollection.Indexes().List(context.TODO())
	//for list.Next(context.TODO()) {
	//	var m bson.M
	//	list.Decode(&m)
	//	fmt.Println(m)
	//}

	//Index 삭제
	//bookCollection.Indexes().DropOne(context.TODO(), "content_text")

	//Index 생성
	model := mongo.IndexModel{Keys: bson.D{{"title", "text"}}}
	name, err := bookCollection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created : ", name)

	filter := bson.D{{"$text", bson.D{{"$search", "article02"}}}}
	cursor, _ := bookCollection.Find(context.TODO(), filter)
	for cursor.Next(context.TODO()) {
		m := bson.M{}
		cursor.Decode(&m)
		fmt.Println(m)
	}
}
