package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GORM/GormTest/gormTest/domain"
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

	////Primary Key
	//db.First(&actor, "actor_id = ?", 10)
	////일반 SQL 문
	//db.Find(&actor, "first_name = ?", "James")
	//db.Find(&film, "rental_duration > ? AND rating = ?", 5, "R")
	////Struct
	//db.Find(&film, domain.Film{Rating: "G"})
	////Map
	//db.Find(&actor, map[string]interface{}{"first_name": "Tom"})

	//출력 포맷
	for _, film := range film {
		fmt.Println(film.Title, "\t", film.Rating, "\t", film.Rental_duration, "\t", film.Rental_rate, "\t", film.Replacement_cost)
	}
	for _, actor := range actor {
		fmt.Println(actor.Actor_id, "\t", actor.First_name, "\t", actor.Last_name)
	}
}
