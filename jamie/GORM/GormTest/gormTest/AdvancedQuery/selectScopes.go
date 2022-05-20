package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"obigo-go-mentoring/jamie/GORM/GormTest/gormTest/domain"
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

	var film []domain.Film
	db.Scopes(RatingAdult, AVGLength).Find(&film)
	log.Println(film)
}

//관람등급 R
func RatingAdult(db *gorm.DB) *gorm.DB {
	return db.Where("rating = ?", "R")
}

//상영시간 평균 이하
func AVGLength(db *gorm.DB) *gorm.DB {
	return db.Where("length > ?", "AVG(length)")
}
