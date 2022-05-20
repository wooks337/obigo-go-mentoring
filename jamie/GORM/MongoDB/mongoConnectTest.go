package main

import "obigo-go-mentoring/jamie/GORM/MongoDB/conn"

func main() {
	//MongoDB 연결
	conn := conn.MongoConnect()
	//사용할 DB와 Collection 지정
	postCollection := conn.Database("Jamie_Go").Collection("posts")

}
