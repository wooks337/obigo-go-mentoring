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
	cursor, err := bookCollection.Find(context.TODO(), bson.D{})

	defer cursor.Close(context.TODO())

	//기본 출력
	for cursor.Next(context.TODO()) {
		//해당 document 를 지금,나중에 사용 가능
		//에러가 발생하지 않음
		//컨텍스트가 만료되지 않음
	}
	if err := cursor.Err(); err != nil {

	}

	//출력2
	for {
		if cursor.TryNext(context.TODO()) { //참
			//해당 document 를 지금 사용 가능
			//에러가 발생하지 않음
			//컨텍스트가 만료되지 않음
			continue
		}

		if err := cursor.Err(); err != nil {

		}
	}

	//출력3
	var results []bson.D
	if err := cursor.All(context.TODO(), &results); err != nil {

	}
	for _, result := range results {
		fmt.Println(result)
	}

}
