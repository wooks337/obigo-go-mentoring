package main

import (
	"fmt"
	_6_gorm "goMod"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	films := []_6_gorm.Film{}

	//업데이트 완료 전까지 다른 조회는 일어나지 않음 (증명 방법?)
	db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&films)
	fmt.Println(films)

}
