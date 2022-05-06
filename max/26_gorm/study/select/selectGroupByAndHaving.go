package main

import (
	"fmt"
	stru "goMod"
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

	res := make([]result, 1)

	db.Model(&stru.Film{}).Select("length, count(rental_rate) as cnt").
		Group("length").Order("cnt desc").Find(&res)

	db.Model(&stru.Film{}).Select("length, count(rental_rate) as cnt").
		Group("length").Having("cnt > ?", 10).Order("cnt desc").Find(&res)

	fmt.Println(res)

	//rows, _ := db.Table("film").Select("length, count(rental_rate) as cnt").
	//	Group("length").Order("cnt desc").Rows()
	//
	//for rows.Next() {
	//	r := result{}
	//	rows.Scan(&r.Length, &r.Cnt)
	//	fmt.Println(r)
	//}

	results := make([]result, 1)
	db.Table("film").Select("length, count(rental_rate) as cnt").
		Group("length").Order("cnt desc").Scan(&results)

	fmt.Println(results)
}

type result struct {
	Length int
	Cnt    int
}
