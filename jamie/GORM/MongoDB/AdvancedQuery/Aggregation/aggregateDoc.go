package main

import (
	"obigo-go-mentoring/jamie/GORM/MongoDB/conn"
)

func main() {
	conn := conn.MongoConnect()
	pizzaCollection := conn.Database("Jamie_Go").Collection("pizza")

	//insert Data

	//docs := []interface{}{
	//	bson.D{{"_id", 0}, {"name", "Pepperoni"}, {"size", "small"}, {"price", 19.05}, {"quantity", 10}, {"toppings", bson.A{"bacon", "garlic sauce", "pineapple"}}},
	//	bson.D{{"_id", 1}, {"name", "Pepperoni"}, {"size", "medium"}, {"price", 20.50}, {"quantity", 20}, {"toppings", bson.A{"paprika", "garlic sauce", "olive, cheese"}}},
	//	bson.D{{"_id", 2}, {"name", "Pepperoni"}, {"size", "large"}, {"price", 21.00}, {"quantity", 30}, {"toppings", bson.A{"bacon", "olive", "pineapple", "cheese"}}},
	//	bson.D{{"_id", 3}, {"name", "Hawaiian"}, {"size", "small"}, {"price", 15.11}, {"quantity", 25}, {"toppings", bson.A{"bacon", "olive", "garlic sauce", "pineapple", "cheese"}}},
	//	bson.D{{"_id", 4}, {"name", "Hawaiian"}, {"size", "medium"}, {"price", 16.30}, {"quantity", 40}, {"toppings", bson.A{"olive", "pineapple", "cheese"}}},
	//	bson.D{{"_id", 5}, {"name", "Cheese"}, {"size", "small"}, {"price", 11.50}, {"quantity", 5}, {"toppings", bson.A{"bacon", "paprika", "olive", "pineapple", "cheese"}}},
	//	bson.D{{"_id", 6}, {"name", "Cheese"}, {"size", "large"}, {"price", 12.15}, {"quantity", 10}, {"toppings", bson.A{"pickle", "hot sauce", "sweetpotato moose", "basil"}}},
	//	bson.D{{"_id", 7}, {"name", "Vegan"}, {"size", "medium"}, {"price", 13.00}, {"quantity", 10}, {"toppings", bson.A{"pickle", "hot sauce", "potato", "basil"}}},
	//}
	//result, _ := pizzaCollection.InsertMany(context.TODO(), docs)
	//fmt.Println(result.InsertedIDs)

	////$match, $project, $sort

	////토핑에 pineapple이 들어가고 _id는 안 보여주고 최저가로 정렬
	//matchStage := bson.D{{"$match", bson.D{{"toppings", "bacon"}}}}
	//projectMatch := bson.D{{"$project", bson.D{{"_id", 0}}}}
	//sortMatch := bson.D{{"$sort", bson.D{{"price", 1}}}}
	//
	//cursor, _ := pizzaCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, projectMatch, sortMatch})
	//var results []bson.M
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Printf("%v 피자의 %v 사이즈 가격은 %v이고 토핑은 %v 이/가 올라갑니다\n", result["name"], result["size"], result["price"], result["toppings"])
	//}

	////$group

	////피자 종류별 평균 가격과 개수 출력
	//groupStage := bson.D{
	//	{"$group", bson.D{
	//		{"_id", "$name"},
	//		{"avgPrice", bson.D{{"$avg", "$price"}}},
	//		{"totalQty", bson.D{{"$sum", "$quantity"}}},
	//	}},
	//}
	//cursor, _ := pizzaCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	//var results []bson.M
	//cursor.All(context.TODO(), &results)
	//for _, result := range results {
	//	fmt.Printf("%v 피자 평균가격: %v\n", result["_id"], result["avgPrice"])
	//	fmt.Printf("%v 피자 개수: %v\n", result["_id"], result["totalQty"])
	//}
}
