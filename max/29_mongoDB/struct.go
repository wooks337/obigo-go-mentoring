package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongoTest/connect"
)

func main() {

	mongoDB, err := connect.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer func() {
		mongoDB.Disconnect(context.TODO())
	}()

	gradeCol := mongoDB.Database("maxSample").Collection("grade")

	////insert
	//addBook := book{
	//	company:  "애플",
	//	title:    "제목",
	//	likes:    150,
	//	writer:   "스티븐",
	//	intarray: nil,
	//}

	//newGrade := grade{
	//	english:   "A",
	//	korean:    "B",
	//	maths:     "C",
	//	science:   "D",
	//	studentId: "05",
	//}
	//
	//res, _ := gradeCol.InsertOne(context.TODO(), newGrade)
	//fmt.Println(res.InsertedID)

	//find
	//var findGrade []grade
	//
	//cursor, _ := gradeCol.Find(context.TODO(), bson.D{})
	//fmt.Println("====")
	//fmt.Println(findGrade)
	//err = cursor.All(context.TODO(), &findGrade)
	//if err != nil {
	//	fmt.Println("err 발생 : ", err)
	//	return
	//}
	//fmt.Println(findGrade)

	fmt.Println("====")
	var findGrade grade
	gradeCol.FindOne(context.TODO(), bson.D{}).Decode(&findGrade)

	fmt.Println(findGrade)
}

type grade struct {
	ID        primitive.ObjectID `bson:"_id"`
	English   string             `bson:"english"`
	Korean    string             `bson:"korean"`
	Maths     string             `bson:"maths"`
	Science   string             `bson:"science"`
	StudentId string             `bson:"student_id"`
}

type book struct {
	ID       primitive.ObjectID
	Title    string
	Company  string
	Likes    int
	Writer   string
	Intarray []int
}
