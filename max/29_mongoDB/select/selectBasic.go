package main

import (
	"context"
	"encoding/json"
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

	//title이 naruto 하나 검색
	//res := bson.M{}
	//err = bookCollection.FindOne(context.TODO(), bson.D{{"title", "Naruto"}}).
	//	Decode(&res)
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		fmt.Println("no data")
	//	} else {
	//		fmt.Println("err : ", err)
	//		return
	//	}
	//}
	//fmt.Println(res)

	//likes 20개 이상출력
	find, err := bookCollection.Find(context.TODO(), bson.D{
		{"likes", bson.D{{"$gte", 20}}},
	})
	for find.Next(context.TODO()) {
		var elem bson.M
		err := find.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		elem2, err := json.MarshalIndent(elem, "", "  ")
		fmt.Printf("%s", elem2)
	}

}
