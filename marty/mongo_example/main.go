package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UID      string             `bson:"id"`
	Username string             `bson:"username"`
	Password string             `bson:"passwrod"`
}

func main() {

	credential := options.Credential{
		Username: "root",
		Password: "root",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	connect, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		err = fmt.Errorf("연결실패 : %v", err)
	}
	defer func() {
		connect.Disconnect(context.TODO())
	}()

	err = connect.Ping(context.TODO(), nil)
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}

	fmt.Println("MongoDB Connect Success")

	userCollection := connect.Database("mongodb_tutorial").Collection("user")
	ctx := context.TODO()

	// GET
	user := new(User)
	err = userCollection.FindOne(ctx, bson.M{}).Decode(user)
	if err != nil {
		fmt.Println("Invalid db data")
	}

	fmt.Println("User : ", user)

	// SET
	user = &User{
		ID:       primitive.ObjectID{},
		UID:      "2",
		Username: "max",
		Password: "1231",
	}
	_, _ = userCollection.InsertOne(ctx, user)
}
