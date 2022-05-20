package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//======TeamServer======
func main() {
	//MongoDB에 Go 연결
	clientOptions := options.Client().ApplyURI("mongodb://10.28.3.180:27017").
		SetAuth(options.Credential{
			AuthSource: "",
			Username:   "root",
			Password:   "root",
		})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made XD")

	//MongoDB 연결 종료
	usersCollection := client.Database("blog").Collection("posts")
	fmt.Println(usersCollection)
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection End Xp")
}
