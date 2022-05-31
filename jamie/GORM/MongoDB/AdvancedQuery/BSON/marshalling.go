package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.D{{"age", 8}}

	var result bson.D
	studentColl.FindOne(context.TODO(), filter).Decode(&result)

	fmt.Println(result)

	////Query
	//var result Student
	//
	//opts := options.FindOne().SetProjection(bson.D{{"_id", 0}, {"tags", 1}})
	//studentColl.FindOne(context.TODO(), bson.D{}, opts).Decode(&result)
	//fmt.Println(result)

}
