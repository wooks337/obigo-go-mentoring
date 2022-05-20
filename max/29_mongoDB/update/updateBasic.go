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

	filter := bson.D{{"title", "Naruto"}}
	updateBson := bson.D{
		{"$set", bson.D{{"writer", "ninza2"}}},
	}

	res, err := bookCollection.UpdateOne(context.TODO(), filter, updateBson)

	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println("Suc")
	}
	fmt.Println(res)
}
