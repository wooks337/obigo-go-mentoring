package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	//==== Data retrieve 시 Options

	//filter := bson.D{{"tags", bson.D{{"$in", bson.A{"HtMl", "MySQL", "Spring"}}}}}
	//coll := &options.Collation{
	//	Locale:    "en_US",
	//	Strength:  1,
	//	CaseLevel: false,
	//}
	//projection := bson.D{{"_id", 0}, {"likes", 1}, {"tags", 1}, {"title", 1}}
	//sort := bson.D{{"likes", 1}}
	//
	//opt := options.Find().SetCollation(coll).SetProjection(projection).SetLimit(3).SetSort(sort)
	//cursor, err := postCollection.Find(context.TODO(), filter, opt)
	//if err != nil {
	//	log.Println("===1")
	//	log.Fatal(err)
	//}
	//var results []bson.D
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	log.Println("===2")
	//	log.Fatal(err)
	//}
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//==== Aggregation을 이용한 Data Retrieve

	//matchStage := bson.D{{"$match", bson.D{{"category", "News"}}}}
	//projectStage := bson.D{{"$project", bson.D{{"_id", 0}, {"comments", 0}, {"date", 0}}}}
	//sortStage := bson.D{{"$sort", bson.D{{"title", -1}}}}
	//limitStage := bson.D{{"$limit", 3}}
	//
	//cursor, err := postCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, projectStage, sortStage, limitStage})
	//if err != nil {
	//	log.Println("===1. err")
	//	log.Fatal(err)
	//}
	//var results []bson.D
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//==== Distinct를 이용한 특정 Data Retrieve

	//results, err := postCollection.Distinct(context.TODO(), "comments", bson.D{{"likes", 4}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//==== Text Search

	////==== create text index
	//model := mongo.IndexModel{Keys: bson.D{{"tags", "text"}}}
	//name, err := postCollection.Indexes().CreateOne(context.TODO(), model)
	//if err != nil {
	//	log.Println("==1. err")
	//	log.Fatal(err)
	//}
	//fmt.Println("생성된 인덱스명: ", name)

	filter := bson.D{{"$text", bson.D{{"$search", "mysql"}}}}
	cursor, err := postCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("===1.err")
		log.Fatal(err)
	}

	var result []bson.D
	cursor.All(context.TODO(), &result)
	for _, results := range result {
		fmt.Println(results)
	}

}
