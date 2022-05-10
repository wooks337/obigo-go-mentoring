package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/sakila?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})

	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	//res := []FilmActors{}
	//db.Table("film").Select("a.actor_id, film.film_id, title, first_name, last_name").
	//	Joins("JOIN film_actor fa ON fa.film_id = film.film_id").
	//	Joins("JOIN actor a ON a.actor_id = fa.actor_id").Scan(&res)
	//
	//fmt.Println(res)

	//rows, _ := db.Table("film").Select("a.actor_id, film.film_id, title, first_name, last_name").
	//	Joins("JOIN film_actor fa ON fa.film_id = film.film_id").
	//	Joins("JOIN actor a ON a.actor_id = fa.actor_id").Rows()
	//for rows.Next() {
	//	f := FilmActors{}
	//	rows.Scan(&f.Actor_id, &f.Film_id, &f.Title, &f.First_name, &f.Last_name)
	//	fmt.Println(f)
	//}

	//res := []_6_gorm.Film{}
	//db.Joins("film_actor").Find(&res)
	//fmt.Println(res)

}

type FilmActors struct {
	Actor_id   int
	Film_id    int
	Title      string
	First_name string
	Last_name  string
}
