package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	//=====Raw=====
	//var films []_6_gorm.Film
	//db.Raw("SELECT film_id, title FROM film WHERE film_id <= ?", 10).Scan(&films)
	//fmt.Println(films)

	//var lengthAvg float64
	//db.Raw("SELECT AVG(length) FROM film").Scan(&lengthAvg)
	//fmt.Println(lengthAvg)

	//db.Raw("UPDATE film SET description = 'none' WHERE film_id = ?", 1017).Scan(&_6_gorm.Film{})

	//=====Exec=====
	//db.Exec("UPDATE film SET description = 'none' WHERE film_id = ?", 1017)

	//db.Exec("UPDATE film SET length = ? WHERE film_id = ?", gorm.Expr("length + ?", 1), 1017)
	//db.Exec("UPDATE film SET length = ? WHERE film_id = (SELECT film_id FROM (?) as sub)", gorm.Expr("length + ?", 1), db.Table("film").Select("film_id").Order("film_id DESC").Limit(1))

}
