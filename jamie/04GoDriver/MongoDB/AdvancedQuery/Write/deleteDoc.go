package main

import (
	"obigo-go-mentoring/jamie/04GoDriver/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	commentCollection := conn.Database("Jamie_Go").Collection("comments")

	////deleteOne()
	//filter := bson.D{{"com_name", "gayle"}}
	//result, _ := commentCollection.DeleteOne(context.TODO(), filter)
	//fmt.Println("삭제된 doc 개수: ", result.DeletedCount)

	////deleteMany()
	//filter := bson.D{{"_id", bson.D{{"$gt", 5}}}}
	//result, _ := commentCollection.DeleteMany(context.TODO(), filter)
	//fmt.Println("삭제된 doc 개수: ", result.DeletedCount)
}
