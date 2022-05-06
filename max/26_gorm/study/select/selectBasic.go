package main

import (
	"fmt"
	"goMod"
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

	//=====단일 검색====
	//var film _6_gorm.Film
	//db.First(&film, 1) //PK로 찾기
	//fmt.Println(film)
	//
	//db.First(&film, "film_id = ?", 2)
	//fmt.Println(film)

	//=====복수 검색====
	films := make([]_6_gorm.Film, 1)
	//res := db.Find(&films, []int{4, 5, 6}) //IN (4,5,6)
	//fmt.Println(films)
	//fmt.Println(res.RowsAffected) //검색 결과출력

	res := db.Find(&films, "rental_rate = ?", "2.99")
	fmt.Println(films)
	fmt.Println(res.RowsAffected) //검색 결과출력

}
