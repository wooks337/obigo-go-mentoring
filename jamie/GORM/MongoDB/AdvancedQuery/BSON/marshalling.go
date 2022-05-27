package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

type Addr struct {
	Street string
	City   string
	State  string
}
type Student struct {
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Addr      Addr   `bson:"inline"`
	Age       int
	Tags      []int `bson:"tags,omitempty"`
}

func main() {
	connect := conn.MongoConnect()
	//addressColl := connect.Database("Schools").Collection("addrs")
	studentColl := connect.Database("Schools").Collection("students")

	////Insert Document
	//address1 := Addr{"1 Lakewood Way", "Elwood City", "PA"}
	//student1 := Student{FirstName: "Arthur", Addr: address1, Age: 8, Tags: []int{12, 13}}
	//_, err1 := addressColl.InsertOne(context.TODO(), address1)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//_, err2 := studentColl.InsertOne(context.TODO(), student1)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}

	//Query
	var result []byte
	var test []int
	filter := bson.D{{"state", "PA"}}
	opts := options.FindOne().SetProjection(bson.D{{"_id", 0}, {"tags", 1}})
	studentColl.FindOne(context.TODO(), filter, opts).Decode(&result)
	bson.Unmarshal(result, &test)
	fmt.Println(test)
	fmt.Println(result)

	//fmt.Println(reflect.TypeOf(result))

	//doc, err := bson.Marshal(bson.D{{"tags", bson.A{24, 55, 32, 10, 13, 85, 8}}})
	//if err != nil {
	//	log.Println(err)
	//}
	//var test bson.D
	//err = bson.Unmarshal(doc, &test)
	//fmt.Println(test)
}
