package main

import (
	"fmt"
	"goMod/study/createTable/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/gormMax?charset=utf8&parseTime=True&loc=Local"
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

	s1 := []domain.Student{}
	s2 := []domain.Student{}

	queryDB := db.Where("ID IN (?)", []int{1, 2, 3, 4, 5, 6, 7, 8})
	queryDB = queryDB.Find(&s1)
	queryDB.Where("age > 23").Find(&s2)

	for _, student := range s1 {
		fmt.Println(student.Name)
	}
	fmt.Println("=================")
	for _, student := range s2 {
		fmt.Println(student.Name)
	}

}
