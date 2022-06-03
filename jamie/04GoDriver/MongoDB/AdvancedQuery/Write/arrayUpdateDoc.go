package main

import (
	"obigo-go-mentoring/jamie/04GoDriver/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	postCollection := conn.Database("Jamie_Go").Collection("posts")

	////FindAndUpdate()
	////$
	//filter := bson.D{{"tags", bson.D{{"$in", bson.A{"java"}}}}}
	//update := bson.D{{"$set", bson.D{{"tags.$", "javac"}}}}
	//opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	//
	//var updatedDoc bson.D
	//err := postCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
	//if err != nil {
	//	log.Println("==err")
	//	log.Fatal(err)
	//}
	//fmt.Println(updatedDoc)

	////$[]
	//identifier := []interface{}{bson.D{{"smaller", bson.D{{"$lt", 10}}}}}
	//update := bson.D{{"$inc", bson.D{{"tags.$[smaller]", 5}}}}
	//opts := options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{Filters: identifier}).SetReturnDocument(options.After)
	//
	//var updatedDoc bson.D
	//err := postCollection.FindOneAndUpdate(context.TODO(), bson.D{}, update, opts).Decode(&updatedDoc)
	//if err != nil {
	//	log.Println("===err")
	//	log.Fatal(err)
	//}
	//fmt.Println(updatedDoc)

	////addToSet
	//filter := bson.D{{"title", "Post15"}}
	//update := bson.D{{"$addToSet", bson.D{{"tags", "kotlin"}}}}
	//
	//result, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	//fmt.Println("수정된 docu 갯수", result.ModifiedCount)
	//
	////addToSet-each
	//filter := bson.D{{"title", "Post15"}}
	//update := bson.D{{"$addToSet", bson.D{{"tags", bson.D{{"$each", bson.A{"postgresql", "oracle"}}}}}}}
	//
	//result, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	//fmt.Println("수정된 docu 갯수", result.ModifiedCount)
	//
	////pop
	//filter := bson.D{{"title", "Post15"}}
	//update := bson.D{{"$pop", bson.D{{"tags", 1}}}}
	//
	//result, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	//fmt.Println("수정된 doc 개수", result.ModifiedCount)
	//
	////push
	//filter := bson.D{{"category", "Obigo"}}
	//update := bson.D{{"$push", bson.D{{"comments", bson.D{{"com_id", 2}, {"recommend", 22}}}}}}
	//
	//result, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	//fmt.Println("수정된 doc 개수", result.ModifiedCount)
	//
	////pull
	//filter := bson.D{{"category", "Obigo"}}
	//update := bson.D{{"$pull", bson.D{{"comments", bson.D{{"com_id", 2}}}}}}
	//
	//result, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	//fmt.Println("수정된 doc 개수", result.ModifiedCount)

}
