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

	distinct, _ := bookCollection.Distinct(context.TODO(), "company",
		bson.D{{"likes", bson.D{{"$gte", 10}}}})
	for _, res := range distinct {
		fmt.Println(res)
	}

}
