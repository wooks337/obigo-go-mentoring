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
		DryRun: true,                                //SQL이 실제 실행 되지 않음
	})

	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	majors := []domain.Major{}
	db.Find(&majors)
	fmt.Println(majors)

	students := []domain.Student{}
	students = append(students, domain.Student{Name: "테스터", Age: 24, MajorId: 1})
	db.Create(&students)

	statement := db.Find(&majors).Statement
	fmt.Println(statement.SQL.String())
}
