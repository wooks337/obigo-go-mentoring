package main

import (
	"obigo-go-mentoring/jamie/04GoDriver/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")
	//commentCollection := conn.Database("Jamie_Go").Collection("comments")

	//===정렬 조회 : SetSort()
	////likes가 2 이상인 모든 쿼리를 likes 개수에 따라 정렬
	//opt := options.Find()
	//opt.SetSort(bson.D{{"likes", -1}})
	//sortCursor, err := postCollection.Find(context.TODO(), bson.D{{"likes", bson.D{{"$gt", 2}}}}, opt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var postSorted []bson.M
	//if err = sortCursor.All(context.TODO(), &postSorted); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(postSorted)

	//===Projection 조회 : SetProjection()
	////like가 2 이상인 document를 comment와 id는 제외하고 보여주되 tags는 두개씩만 조회
	//opt := options.Find().SetProjection(bson.D{{"comments", 0}, {"_id", 0}, {"tags", bson.D{{"$slice", 2}}}})
	//cursor, _ := postCollection.Find(context.TODO(), bson.D{{"likes", bson.D{{"$gt", 2}}}}, opt)
	//
	//var result []bson.D
	//cursor.All(context.TODO(), &result)
	//for _, res := range result {
	//	fmt.Println(res)
	//}

	//===논리 연산자 $or : 다중 조건 만족 docu 조회
	////projection으로 likes랑 recommend만 출력
	//filter := bson.D{{"$or", bson.A{
	//	bson.D{{"likes", bson.D{{"$gt", 1}, {"$lte", 5}}}},
	//	bson.D{{"comments.recommend", bson.D{{"$gte", 10}}}}}}}
	//opt := options.Find().SetProjection(bson.D{{"_id", 0}, {"likes", 1}, {"comments", 1}})
	//
	//cursor, _ := postCollection.Find(context.TODO(), filter, opt)
	//var results []bson.D
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//===문자열 연산자 $regex
	//filter := bson.D{{"com_text", bson.D{{"$regex", "^p"}}}}
	//cursor, _ := commentCollection.Find(context.TODO(), filter)
	//var results []bson.D
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//===배열 연산자 $all
	//filter := bson.D{{"tags", bson.D{{"$all", bson.A{"java"}}}}}
	//opt := options.Find().SetProjection(bson.D{{"_id", 0}, {"likes", 1}, {"comments", 1}, {"tags", 1}})
	//
	//cursor, _ := postCollection.Find(context.TODO(), filter, opt)
	//
	//var results []bson.D
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//===배열 연산자 $elemMatch
	//filter := bson.D{{"comments", bson.D{{"$elemMatch", bson.D{{"com_id", 1}}}}}}
	//opt := options.Find().SetProjection(bson.D{{"_id", 0}, {"likes", 1}, {"comments", 1}, {"tags", 1}})
	//
	//cursor, _ := postCollection.Find(context.TODO(), filter, opt)
	//var results []bson.D
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Println(result)
	//}

}
