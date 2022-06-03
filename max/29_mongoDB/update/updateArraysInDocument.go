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

	////pushArrays
	//filter := bson.D{{"company", "기아"}}
	//pushBson := bson.D{{"$push", bson.D{{"intarray", bson.D{{"$each", []int{8, 9}}}}}}}
	//res, _ := bookCollection.UpdateMany(context.TODO(), filter, pushBson)
	//
	//fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)

	////pullArrays
	//filter := bson.D{{"company", "기아"}}
	//pullBson := bson.D{{"$pull", bson.D{{"intarray", bson.D{{"$in", []int{8, 9}}}}}}}
	//res, _ := bookCollection.UpdateMany(context.TODO(), filter, pullBson)
	//
	//fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)

	////inc
	//filter := bson.D{{"company", "기아"}, {"intarray", bson.D{{"$exists", true}}}}
	////incLikesBson := bson.D{{"$inc", bson.D{{"likes", 1}}}}
	//incArrayBson := bson.D{{"$inc", bson.D{{"intarray.$", 1}}}}
	//incArrayBson2 := bson.D{{"$inc", bson.D{{"intarray.$[]", 1}}}}
	//
	//res, _ := bookCollection.UpdateMany(context.TODO(), filter, incArrayBson)
	//fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)

	//특정 Array
	filter := bson.D{{"company", "기아"}, {"intarray", bson.D{{"$exists", true}}}}
	identifier := []interface{}{bson.D{{"smaller", bson.D{{"$lte", 2}}}}}
	update := bson.D{{"$inc", bson.D{{"intarray.$[smaller]", 10}}}}
	opts := options.Update().SetArrayFilters(options.ArrayFilters{Filters: identifier})

	res, _ := bookCollection.UpdateMany(context.TODO(), filter, update, opts)
	fmt.Println("Suc. match : ", res.MatchedCount, ", modified : ", res.ModifiedCount)
}
