package connect

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	credential := options.Credential{
		Username: "root",
		Password: "root",
	}

	clientOptions := options.Client().ApplyURI("mongodb://10.28.3.180:27017").SetAuth(credential)
	connect, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
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
}
