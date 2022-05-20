package conn

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func MongoConnect() (client *mongo.Client) {
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
