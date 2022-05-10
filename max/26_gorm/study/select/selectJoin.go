package main

import (
	"fmt"
	_6_gorm "goMod"
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

	//res := []filmActors{}
	//db.Table("film").Select("a.actor_id, film.film_id, title, first_name, last_name").
	//	Joins("JOIN film_actor fa ON fa.film_id = film.film_id").
	//	Joins("JOIN actor a ON a.actor_id = fa.actor_id").Scan(&res)
	//
	//fmt.Println(res)

	//rows, _ := db.Table("film").Select("a.actor_id, film.film_id, title, first_name, last_name").
	//	Joins("JOIN film_actor fa ON fa.film_id = film.film_id").
	//	Joins("JOIN actor a ON a.actor_id = fa.actor_id").Rows()
	//for rows.Next() {
	//	f := filmActors{}
	//	rows.Scan(&f.actor_id, &f.film_id, &f.title, &f.first_name, &f.last_name)
	//	fmt.Println(f)
	//}

	res := []_6_gorm.Film{}
	db.Joins("film_actor").Find(&res)
	fmt.Println(res)

}

type filmActors struct {
	actor_id   int
	film_id    int
	title      string
	first_name string
	last_name  string
}
