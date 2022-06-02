package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
	"time"
)

func main() {
	//MongoDB 연결
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")
	//commentCollection := conn.Database("Jamie_Go").Collection("comments")

	////===== Insert =====////

	//===InsertOne

	post1 := bson.D{
		{"title", "Post15"},
		{"category", "News"},
		{"likes", 1},
		{"body", "Body of post."},
		{"comments", bson.A{
			bson.D{
				{"com_id", 1},
				{"recommend", 0},
			},
			bson.D{
				{"com_id", 2},
				{"recommend", 0},
			},
		},
		},
		{"date", time.Stamp},
		{"tags", bson.A{"java", "spring", "backend", "webserver"}},
	}
	result, err := postCollection.InsertOne(context.TODO(), post1)
	if err != nil {
		fmt.Println("err!!", err)
	}
	fmt.Println("==insert 결과==", result.InsertedID)

	//===InsertMany

	//posts := []interface{}{
	//	bson.D{
	//		{"author", "Liz"},
	//		{"category", "AI"},
	//		{"views", 3},
	//		{"content", "this is posting 2"},
	//	},
	//	bson.D{
	//		{"author", "Mela"},
	//		{"category", "HR"},
	//		{"views", 5},
	//		{"content", "this is posting 3"},
	//	},
	//}
	//results, err := postCollection.InsertMany(context.TODO(), posts)
	//if err != nil {
	//	fmt.Println("InsertMany 오류")
	//	log.Fatal(err)
	//}
	//fmt.Println(results.InsertedIDs)

	//===Insert 시 Bson 다른거 써보기

	////===배열 넣기
	//post2 := bson.D{
	//	{"author", "Jeremy"},
	//	{"category", "language"},
	//	{"views", 54},
	//	{"tags", bson.A{"francais", "espagnol", "coreen", "arabe"}},
	//	{"content", "This is language exchange post"},
	//}
	//result, err := postCollection.InsertOne(context.TODO(), post2)
	//if err != nil {
	//	fmt.Println("err!!", err)
	//}
	//fmt.Println("==insert 결과==", result.InsertedID)
}
