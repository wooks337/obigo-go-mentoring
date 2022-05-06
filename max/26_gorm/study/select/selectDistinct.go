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

	res := make([]result2, 1)

	find := db.Model(&stru.Film{}).Select("length, count(rental_rate) as cnt").
		Group("length").Order("cnt desc").Find(&res)
	//fmt.Println(res)
	fmt.Println(find.RowsAffected)

	find2 := db.Model(&stru.Film{}).Distinct("length").Select("length").Find(&res)
	fmt.Println(res)
	fmt.Println(find2.RowsAffected)
}

type result2 struct {
	Length int
	Cnt    int
}
