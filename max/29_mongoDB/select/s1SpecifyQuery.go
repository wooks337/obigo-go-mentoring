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

	//All 메서드 이용하여 Slice에 담기
	//var results []bson.M
	//cursor, _ := bookCollection.Find(context.TODO(), bson.D{})
	//
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	//fmt.Println(result)
	//	indent, _ := json.MarshalIndent(result, "", " ")
	//	fmt.Printf("%s\n\n", indent)
	//	//fmt.Println(result["_id"])
	//}

	//$and연산자 사용
	//filter := bson.D{
	//	{"$and", bson.A{ //A사용
	//		bson.D{{"likes", bson.D{{"$gte", 15}}}},
	//		bson.D{{"title", bson.D{{"$regex", "article"}}}}, //%article%
	//		//bson.D{{"title", bson.D{{"$regex", "^article"}}}}, //%article
	//		//bson.D{{"title", bson.D{{"$regex", "02$"}}}}, //%02
	//	}},
	//}
	//cursor2, _ := bookCollection.Find(context.TODO(), filter)
	//for cursor2.Next(context.TODO()) {
	//	m := bson.M{}
	//	cursor2.Decode(&m)
	//	fmt.Println(m["title"], ", ", m["likes"])
	//}

	//$exists 사용
	//cursor3, _ := bookCollection.Find(context.TODO(), bson.D{
	//	{"comments", bson.D{{"$exists", false}}},
	//})
	//for cursor3.Next(context.TODO()) {
	//	m := bson.M{}
	//	cursor3.Decode(&m)
	//	fmt.Println(m)
	//}

	//$all 사용, array
	cursor4, _ := bookCollection.Find(context.TODO(), bson.D{
		{"intarray", bson.D{{"$all", bson.A{1, 2}}}},
	})
	for cursor4.Next(context.TODO()) {
		m := bson.M{}
		cursor4.Decode(&m)
		fmt.Println(m)
	}
}
