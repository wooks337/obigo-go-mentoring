package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
	"time"
)

func mongoConn() (client *mongo.Client) {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}
	clientOptions := options.Client().ApplyURI("mongodb://10.28.3.180:27017").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("연결 안됨")
		log.Fatal(err)
	}
	//데이터베이스 커넥션이 너무 길어지는 것을 방지하기 위해 컨텍스트로 생명주기 제어
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	// 연결 확인
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("연결 확인 부분 이상함")
		log.Fatal(err)
	}
	fmt.Println("MongoDB 연결~~~")
	return client
}

func main() {
	//MongoDB 연결
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	////===== Delete =====////

	//DeleteOne

	//첫번째 document 삭제
	result, err := postCollection.DeleteOne(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("document 삭제 error!!")
		log.Fatal(err)
	}
	fmt.Println("첫번째 document 삭제")
	fmt.Println("삭제된 document 수: ", result.DeletedCount)

	//DeleteMany

	////조건: views가 3 이상인 documents
	//filter := bson.D{
	//	{"views", bson.D{{"$gte", 3}}},
	//}
	//results, err := postCollection.DeleteMany(context.TODO(), filter)
	//if err != nil {
	//	fmt.Println("document 삭제 error!!!")
	//	log.Fatal(err)
	//}
	//fmt.Println("조건에 맞는 document 모두 삭제")
	//fmt.Println("삭제된 document 수: ", results.DeletedCount)
}
