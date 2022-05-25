package main

import (
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	commentCollection := conn.Database("Jamie_Go").Collection("comments")

	////insertOne()
	//com := bson.D{
	//	{"_id", 7},
	//	{"com_createdAt", time.Stamp},
	//	{"com_name", "charlie"},
	//	{"com_text", "awesome and perfect"},
	//}
	//cursor, err := commentCollection.InsertOne(context.TODO(), com)
	//if err != nil {
	//	log.Println("==err1")
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(cursor.InsertedID)

	////insertMany()
	//coms := []interface{}{
	//	bson.D{{"_id", 8}, {"com_createdAt", time.Date}, {"com_name", "martine"}, {"com_text", "magnifique and hello"}},
	//	bson.D{{"_id", 9}, {"com_createdAt", time.Stamp}, {"com_name", "gayle"}, {"com_text", "great and nah"}},
	//}
	//
	//result, _ := commentCollection.InsertMany(context.TODO(), coms)
	//fmt.Println("생성된 docu:", len(result.InsertedIDs))
	//for _, id := range result.InsertedIDs {
	//	fmt.Println("생성된 docu id", id)
	//}
}
