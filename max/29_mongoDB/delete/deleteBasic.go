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

	////DeleteOne
	//res, err := bookCollection.DeleteOne(context.TODO(), bson.D{
	//	{"title", "Naruto"},
	//})
	//
	//if err != nil {
	//	fmt.Println("error : ", err)
	//} else {
	//	fmt.Println("Suc")
	//}
	//fmt.Println(res)

	//DeleteMany
	filter := bson.D{{"company", bson.D{{"$exists", false}}}}
	res, err := bookCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println("Suc ", res.DeletedCount)
	}
}
