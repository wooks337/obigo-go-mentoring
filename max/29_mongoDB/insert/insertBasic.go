package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	////InsertOne
	//res, err := bookCollection.InsertOne(context.TODO(), bson.D{
	//	{"company", "대우"}, {"content", "dww"}, {"intarray", []int{1, 2, 3}},
	//	{"likes", 24}, {"title", "DW"}, {"writer", "Son"},
	//})
	//if err != nil {
	//	fmt.Println("error : ", err)
	//} else {
	//	fmt.Println("Suc")
	//}
	//fmt.Println(res.InsertedID)

	////InsertMany
	//docs := []interface{}{
	//	bson.D{{"company", "현대"}, {"content", "article 06"}, {"intarray", []int{3, 2}},
	//		{"likes", 11}, {"title", "article06"}, {"writer", "Lee"}},
	//	bson.D{{"company", "기아"}, {"content", "article 07"}, {"intarray", []int{1, 2}},
	//		{"likes", 11}, {"title", "article07"}, {"writer", "Kim"}},
	//	bson.D{{"company", "현대"}, {"content", "article 08"}, {"intarray", []int{}},
	//		{"likes", 12}, {"title", "article08"}, {"writer", "Sim"}},
	//}
	//res2, err := bookCollection.InsertMany(context.TODO(), docs)
	//if err != nil {
	//	fmt.Println("error : ", err)
	//} else {
	//	fmt.Println("Suc")
	//}
	//fmt.Println(res2)

	//InsertMany Option
	docs2 := []interface{}{
		bson.D{{"_id", 1}, {"country", "Tanzania"}},
		bson.D{{"_id", 2}, {"country", "Lithuania"}},
		bson.D{{"_id", 1}, {"country", "Vietnam"}},
		bson.D{{"_id", 3}, {"country", "Argentina"}},
	}
	opts := options.InsertMany().
		SetBypassDocumentValidation(true). //true 일경우 유효성 검사 생략
		SetOrdered(false)                  //true 일경우 document를 순서대로 서버에 보냄, 오류 발생시 나머지 중단
	res3, err := bookCollection.InsertMany(context.TODO(), docs2, opts)
	list_ids := res3.InsertedIDs
	if err != nil {
		fmt.Printf("A bulk write error occurred, but %v documents were still inserted.\n", len(list_ids))
	}
	for _, id := range list_ids {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}
