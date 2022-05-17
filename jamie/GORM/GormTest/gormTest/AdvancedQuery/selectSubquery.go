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

	////서브 쿼리
	var film domain.Film
	//db.Where("rental_duration > (?)", db.Table("film").Select("AVG(rental_duration)")).Take(&film)
	//log.Println(film)

	////서브쿼리 변수로 사용
	//subQuery := db.Select("AVG(rental_rate)").Table("film")
	//db.Where("rental_rate > (?)", subQuery).Find(&film)
	//log.Println(film)

	////From SubQuery
	db.Table("(?)", db.Model(&film).Select("title", "rental_rate")).Where("rental_rate = ?", 0.99).Take(&film)
	log.Println(film)

}
