package main

import (
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	//===CountDocuments()
	//filter := bson.D{{"likes", bson.D{{"$lt", 20}}}}
	//
	//count, _ := postCollection.CountDocuments(context.TODO(), filter)
	//fmt.Printf("좋아요가 20보다 작은 document %d\n", count)

	//===EstimatedDocumentCount()
	//count, _ := postCollection.EstimatedDocumentCount(context.TODO())
	//fmt.Printf("컬렉션 내 document 개수 %d\n", count)

	//===$count -- Aggregation
	//matchStage := bson.D{{"$match", bson.D{{"likes", bson.D{{"$gt", 1}}}}}}
	//countStage := bson.D{{"$count", "total num"}}
	//
	//cursor, _ := postCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, countStage})
	//var results []bson.D
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Println(result)
	//}
}
