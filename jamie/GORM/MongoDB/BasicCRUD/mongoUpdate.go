package main

import (
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

func main() {
	//MongoDB 연결
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	////===== Update =====////

	////UpdateOne
	//
	////조건1: title이 post1인 document
	//filter1 := bson.D{
	//	{"title", "Post1"},
	//}
	////수정 내용1: category를 Obigo로 수정
	//update1 := bson.D{
	//	{"$set", bson.D{{"category", "Obigo"}}},
	//}
	////updateOne function 선언
	//result1, err := postCollection.UpdateOne(context.TODO(), filter1, update1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("수정된 레코드 개수: ", result1.ModifiedCount)

	//UpdateMany

	////조건2: views가 1보다 큰 document
	//filter2 := bson.D{
	//	{"views", bson.D{{"$gte", 1}}},
	//}
	////수정 내용2: content를 "조회수가 최소 1임"으로 변경
	//update2 := bson.D{
	//	{"$set", bson.D{{"content", "조회수 최소 1임"}}},
	//}
	////updateMany function 선언
	//result2, err := postCollection.UpdateMany(context.TODO(), filter2, update2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("수정된 레코드 개수: ", result2.ModifiedCount)

}
