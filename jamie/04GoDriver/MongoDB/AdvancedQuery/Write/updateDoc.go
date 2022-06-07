package main

import (
	"obigo-go-mentoring/jamie/04GoDriver/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	////UpdateByID()
	//docId, _ := primitive.ObjectIDFromHex("628af19f81a81751ef2e8f89")
	//update := bson.D{{"$currentDate", bson.D{{"date", true}}}}
	//
	//result, _ := postCollection.UpdateByID(context.TODO(), docId, update)
	//fmt.Println("조건에 일치하는 doc: ", result.MatchedCount)
	//fmt.Println("수정된 doc: ", result.ModifiedCount)

	////UpdateOne
	//filter := bson.D{{"title", "Obigo|Home"}}
	//update := bson.D{{"$set", bson.D{{"category", "Obigo"}, {"comments",
	//	bson.D{{"com_id", 1}, {"recommend", 15}}}}}}
	//
	//result, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	//fmt.Println("조건에 일치하는 doc: ", result.MatchedCount)
	//fmt.Println("수정된 doc: ", result.ModifiedCount)

	////ReplaceOne()
	//filter := bson.D{{"category", "Obigo"}}
	//replacement := bson.D{{"body", "Obigo welcome page"}, {"title", "Obigo|Home"}}
	//
	//result, _ := postCollection.ReplaceOne(context.TODO(), filter, replacement)
	//fmt.Println("조건에 일치하는 doc: ", result.MatchedCount)
	//fmt.Println("교체된 doc: ", result.ModifiedCount)

}
