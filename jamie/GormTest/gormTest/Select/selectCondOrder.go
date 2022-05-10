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
	//쿼리 결과 담을 슬라이스
	var film = make([]domain.Film, 1)
	var actor = make([]domain.Actor, 1)

	////SELECT * FROM film ORDER BY rating desc;
	//db.Order("rating desc").Find(&film)

	////다중 정렬
	////SELECT * FROM film ORDER BY rating desc, rental_duration;
	//db.Order("rating desc").Order("rental_duration").Find(&film)
	//db.Order("rating desc, rental_duration").Find(&film)

	////Clauses - ORDER BY FIELD()
	////특정 값 우선정렬
	//db.Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "FIELD(film_id, ?) desc", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true}}).Find(&film)

	//출력 포맷
	for _, film := range film {
		fmt.Println(film.Film_id, "\t", film.Title, "\t", film.Rating, "\t", film.Length, "\t", film.Original_language_id, "\t", film.Rental_duration, "\t", film.Rental_rate, "\t", film.Replacement_cost)
	}
	for _, actor := range actor {
		fmt.Println(actor.Actor_id, "\t", actor.First_name, "\t", actor.Last_name)
	}
}
