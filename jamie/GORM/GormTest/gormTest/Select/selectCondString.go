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

	var film = make([]domain.Film, 1)

	//db.Where("title = ?", "SCHOOL JACKET").Find(&film)
	//db.Where("rental_rate > ?", 4).Limit(5).Find(&film)
	//db.Where("rating IN ?", []string{"G", "PG"}).Limit(5).Find(&film)
	//db.Where("title LIKE ?", "%Agent%").Find(&film)
	//db.Where("rental_rate = ? AND rating = ?", 4.99, "PG").Find(&film)
	db.Where("replacement_cost BETWEEN ? AND ?", 5, 15).Find(&film)

	for _, film := range film {
		fmt.Println(film.Title, "\t", film.Rating, "\t", film.Rental_rate, "\t", film.Replacement_cost)
	}
}
