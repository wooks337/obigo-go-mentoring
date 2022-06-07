package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"obigo-go-mentoring/jamie/03GORM/GormTest/gormTest/domain"
)

//func main() {
//	//======localhost======
//	dsn := "root:jamiekim@(localhost:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		panic(err)
//	}

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	var actor domain.Actor

	//=======First, Take, Last 메소드=======//

	//Take()
	db.Take(&actor) //SELECT * FROM actor LIMIT 1;
	log.Println(actor)

	//First()
	db.First(&actor) //actor_id 정렬 첫번째 레코드 조회
	log.Println(actor)

	db.First(&actor, 10) //actor_id = 10인 레코드 조회
	log.Println(actor)

	db.First(&actor, "Last_name = ?", "GUINESS") //last_name = Guiness인 레코드 조회
	log.Println(actor)

	//Last()
	db.Last(&actor) //actor_id 정렬 마지막 데이터 조회
	log.Println(actor)

	//db.Model()
	result := map[string]interface{}{}
	db.Model(&domain.Actor{}).Take(&result) //// SELECT * FROM `actor` ORDER BY `users`.`id` LIMIT 1
	log.Println(result)

	result2 := map[string]interface{}{}
	db.Table("film").Take(&result2) //// SELECT * FROM `actor` ORDER BY `users`.`id` LIMIT 1
	log.Println(result2)

}
