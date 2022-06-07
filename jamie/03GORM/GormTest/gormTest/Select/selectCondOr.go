package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	//쿼리 결과 담을 슬라이스
	var film = make([]domain.Film, 1)
	var actor = make([]domain.Actor, 1)

	////Or
	//db.Where("first_name = ?", "Julia").Or("last_name = ?", "Hope").Find(&actor)

	////Struct
	//db.Where("title = ?", "SIDE ARK").Or("rating = ?", "PG-13").Find(&film)

	////Map
	//db.Where("actor_id = ?", 55).Or(map[string]interface{}{"first_name": "Ben", "last_name": "Harris"}).Find(&actor)

	//출력 포맷
	for _, film := range film {
		fmt.Println(film.Title, "\t", film.Rating, "\t", film.Rental_duration, "\t", film.Rental_rate, "\t", film.Replacement_cost)
	}
	for _, actor := range actor {
		fmt.Println(actor.Actor_id, "\t", actor.First_name, "\t", actor.Last_name)
	}
}
