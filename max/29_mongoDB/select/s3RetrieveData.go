package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	//Option - sort, Limit, skip, projection
	filter := bson.D{{"likes", bson.D{{"$gte", 13}}}}
	projection := bson.D{{"_id", 0}, {"likes", 1}, {"title", 1}}
	sort := bson.D{{"likes", 1}, {"title", -1}}
	opts := options.Find().SetSort(sort).SetSkip(1).SetLimit(2).SetProjection(projection)
	cursor, _ := bookCollection.Find(context.TODO(), filter, opts)
	for cursor.Next(context.TODO()) {
		m := bson.M{}
		cursor.Decode(&m)
		fmt.Println(m)
	}

	//aggregation pipline - sort, Limit, skip, project
	filter2 := bson.D{{"$match", bson.D{{"likes", bson.D{{"$gte", 13}}}}}}
	projection2 := bson.D{{"$project", bson.D{{"_id", 0}, {"likes", 1}, {"title", 1}}}}
	sort2 := bson.D{{"$sort", bson.D{{"likes", 1}, {"title", -1}}}}
	skip2 := bson.D{{"$skip", 1}}
	limit2 := bson.D{{"$limit", 2}}
	cursor2, _ := bookCollection.Aggregate(context.TODO(), mongo.Pipeline{filter2, projection2, sort2, skip2, limit2})
	for cursor2.Next(context.TODO()) {
		m := bson.M{}
		cursor2.Decode(&m)
		fmt.Println(m)
	}

	//aggregation pipeline - group
	//groupStage := bson.D{
	//	{"$group", bson.D{
	//		{"_id", "$company"},
	//		{"likeAvg", bson.D{{"$avg", "$likes"}}},
	//		{"likeSum", bson.D{{"$sum", "$likes"}}},
	//		{"cnt", bson.D{{"$sum", 1}}},
	//	}},
	//}
	//cursor3, _ := bookCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	//for cursor3.Next(context.TODO()) {
	//	m := bson.M{}
	//	cursor3.Decode(&m)
	//	fmt.Println(m)
	//}

}
