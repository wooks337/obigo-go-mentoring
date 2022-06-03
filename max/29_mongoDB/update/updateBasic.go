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

	//UpdateOne
	//filter := bson.D{{"title", "article08"}}
	//updateBson := bson.D{
	//	{"$set", bson.D{{"writer", "ninza3"}}},
	//}
	//res, err := bookCollection.UpdateOne(context.TODO(), filter, updateBson)
	//
	//if err != nil {
	//	fmt.Println("error : ", err)
	//} else {
	//	fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)
	//}

	//UpdateMany
	filter := bson.D{{"title", bson.D{{"$in", []string{"naruto", "article06", "article07"}}}}}
	updateBson := bson.D{
		{"$push", bson.D{{"comments", bson.D{{"name", "Bravo"}, {"message", "wow"}}}}},
	}
	res, err := bookCollection.UpdateMany(context.TODO(), filter, updateBson)

	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)
	}

	////ReplaceOne
	//filter := bson.D{{"title", "article08"}}
	//updateBson := bson.D{
	//	{"writer", "ninza3"},
	//}
	//res, err := bookCollection.ReplaceOne(context.TODO(), filter, updateBson)
	//
	//if err != nil {
	//	fmt.Println("error : ", err)
	//} else {
	//	fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)
	//}

}
