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
	type Result struct {
		domain.Actor
		domain.Film_actor
	}
	result := make([]Result, 1)

	//// SELECT a.actor_id, fa.film_id
	//// FROM actor as a left join film_actor as fa on fa.actor_id = a.actor _id
	//// GROUP BY a.actor_id, fa.film_id LIMIT 5

	//db.Table("actor as a").Select("a.actor_id, fa.film_id").
	//	Joins("left join film_actor as fa on fa.actor_id = a.actor_id").
	//	Group("a.actor_id, fa.film_id").
	//	Limit(5).
	//	Find(&result)

	////SELECT actor.actor_id, fa.film_id
	//// FROM `actor` left join film_actor as fa on actor.actor_id = fa.actor_id
	//// GROUP BY actor.actor_id, fa.film_id LIMIT 10

	//db.Model(&domain.Actor{}).
	//	Select("actor.actor_id, fa.film_id").
	//	Joins("left join film_actor as fa on actor.actor_id = fa.actor_id").
	//	Group("actor.actor_id, fa.film_id").
	//	Limit(10).
	//	Find(&result)

	for _, result := range result {
		fmt.Println(result.Actor.Actor_id, "\t", result.Film_id)
	}

}
