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

	//count 반환
	//count, _ := bookCollection.CountDocuments(context.TODO(),
	//	bson.D{{"likes", bson.D{{"$gte", 15}}}})
	//fmt.Println(count)

	//Aggregation $count
	//matchStage := bson.D{{"$match", bson.D{{"likes", bson.D{{"$gte", 15}}}}}}
	//countStage := bson.D{{"$count", "totalCount"}}
	//cursor, _ := bookCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, countStage})
	//var results []bson.D
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	panic(err)
	//}
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//EstimatedDocumentCount
	//count, _ := bookCollection.EstimatedDocumentCount(context.TODO())
	//fmt.Println(count)

	//cursor의 갯수
	cursor2, err := bookCollection.Find(context.TODO(), bson.D{})
	fmt.Println(cursor2.RemainingBatchLength())
}
