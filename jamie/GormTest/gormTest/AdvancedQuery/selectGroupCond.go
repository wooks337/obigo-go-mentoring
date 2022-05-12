package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GormTest/gormTest/domain"
)

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
	db.Where(
		db.Where("rating = ?", "G").Where(db.Where("rental_rate = ?", 2.99).Or("rental_rate = ?", 0.99)),
	).Or(
		db.Where("rating = ?", "R").Where("rental_rate = ?", 4.99),
	).Find(&film)

	for _, film := range film {
		fmt.Println(film.Title, "\t", film.Rating, "\t", film.Rental_rate, "\t", film.Replacement_cost)
	}
}
