package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
	"time"
)

func main() {
	conn := conn.MongoConnect()
	commentCollection := conn.Database("Jamie_Go").Collection("comments")

	filter := bson.D{{"com_name", "willy"}}
	update := bson.D{{"$set", bson.D{{"_id", 7}, {"com_createdAt", time.Local}, {"com_text", "hello sweety"}}}}
	opts := options.Update().SetUpsert(true)

	result, _ := commentCollection.UpdateOne(context.TODO(), filter, update, opts)
	fmt.Println("수정된 doc 개수", result.ModifiedCount)
	fmt.Println("upsert된 doc 개수", result.UpsertedCount)

}
