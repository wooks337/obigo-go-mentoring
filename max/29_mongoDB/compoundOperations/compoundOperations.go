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
	filter := bson.D{{"writer", "Son"}}

	//복합 작업 기간 동안 수정 중인 문서에 쓰기 잠금을 설정하여 독립성 유지

	////FindAndUpdate
	update1 := bson.D{{"$set", bson.D{{"title", "DW1"}}}}
	update2 := bson.D{{"$set", bson.D{{"title", "DW2"}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updateDoc bson.D
	bookCollection.FindOneAndUpdate(context.TODO(), filter, update1).Decode(&updateDoc)
	fmt.Println(updateDoc)

	bookCollection.FindOneAndUpdate(context.TODO(), filter, update2, opts).Decode(&updateDoc)
	fmt.Println(updateDoc)

	//FindAndDelete
	var deleteDoc bson.D
	bookCollection.FindOneAndDelete(context.TODO(), filter).Decode(&deleteDoc)
	fmt.Println(deleteDoc)

	//FindAndReplace
	replace := bson.D{{"$set", bson.D{{"title", "DW1"}}}}
	var replaceDoc bson.D
	bookCollection.FindOneAndReplace(context.TODO(), filter, replace).Decode(&replaceDoc)
}
