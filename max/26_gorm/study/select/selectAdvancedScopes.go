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

	//films := []_6_gorm.Film{}
	//db.Scopes(lengthGreaterThan120).Find(&films)
	//fmt.Println(films)

	films2 := []_6_gorm.Film{}
	db.Scopes(lengthGreaterThan120, ratingCheck([]string{"G", "PG"})).Find(&films2)
	fmt.Println(films2)
}
func lengthGreaterThan120(db *gorm.DB) *gorm.DB {
	return db.Where("length > ?", 120)
}
func ratingCheck(status []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("rating IN (?)", status)
	}
}
