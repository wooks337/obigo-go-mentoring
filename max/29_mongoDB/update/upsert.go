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
	filter := bson.D{{"title", "upsertTest"}}
	updateBson := bson.D{
		{"$push", bson.D{{"comments", bson.D{{"name", "Bravo"}, {"message", "wow"}}}}},
		{"$set", bson.D{{"likes", 21}}},
	}
	opts := options.Update().SetUpsert(true)
	res, err := bookCollection.UpdateMany(context.TODO(), filter, updateBson, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount,
		", upserted : ", res.UpsertedCount)

}
