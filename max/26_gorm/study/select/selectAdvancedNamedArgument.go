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

	//한 개
	//films := []_6_gorm.Film{}
	//db.Where("title Like @name1", sql.Named("name1", "%love%")).Find(&films)
	//fmt.Println(films)

	//여러개
	films2 := []_6_gorm.Film{}
	db.Select("title").Where("title Like @name1 OR title Like @name2", map[string]interface{}{
		"name1": "%love%",
		"name2": "%man%",
	}).Find(&films2)
	fmt.Println(films2)

}
