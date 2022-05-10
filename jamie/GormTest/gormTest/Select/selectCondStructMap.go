package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GormTest/gormTest/domain"
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

	var film = make([]domain.Film, 1)

	////Struct
	//db.Where(&domain.Film{Rental_rate: 2.99, Rating: "PG-13"}).Find(&film)
	////특정 필드 검색
	//db.Where(&domain.Film{Rating: "PG", Rental_rate: 2.99}, "rating").Find(&film)

	////Map
	//db.Where(map[string]interface{}{"rating": "NC-17", "replacement_cost": 18.99}).Find(&film)

	for _, film := range film {
		fmt.Println(film.Title, "\t", film.Rating, "\t", film.Rental_rate, "\t", film.Replacement_cost)
	}
}
