package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

func main() {
	//MongoDB 연결
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	////===== Select =====////

	//Find()

	////조건 :view가 1보다 큰 document
	//filter := bson.D{
	//	{"views", bson.D{{"$gt", 1}}},
	//}
	//var results []bson.M
	//cursor, err := postCollection.Find(context.TODO(), filter) //filter 대신 bson.D{}넣으면 조건 없이 전체 조회
	//if err != nil {
	//	fmt.Println("Find 오류")
	//	log.Fatal(err)
	//}
	////조회 결과를 bson으로 바꾸기
	////All 메서드는 cursur를 반복하고 result로 Decoding 한다
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	fmt.Println("조회 결과 bson으로 변환 시 오류")
	//	log.Fatal(err)
	//}
	//fmt.Println("=====조회 결과 모두 출력=====")
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//FindOne()

	//조건 :view가 1보다 큰 document
	filter := bson.D{
		{"views", bson.D{{"$gte", 1}}},
	}
	var result bson.M
	if err := postCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		fmt.Println("FindOne 오류")
		log.Fatal(err)
	}
	fmt.Println("첫번째 조회 결과 출력")
	fmt.Println(result)

}
