package connect

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}

	clientOptions := options.Client().ApplyURI("mongodb://10.28.3.180:27017").SetAuth(credential)
	db, err := mongo.Connect(context.TODO(), clientOptions)
	return db, err
}
