package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"school/domain"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/schoolMax?charset=utf8&parseTime=True&loc=Local"
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

	if err := db.AutoMigrate(&domain.MajorDepartment{}); err != nil {
		fmt.Println("MajorDepartment Err")
	} else {
		fmt.Println("MajorDepartment Suc")
	}

	if err := db.AutoMigrate(&domain.Major{}); err != nil {
		fmt.Println("Major Err")
	} else {
		fmt.Println("Major Suc")
	}

	if err := db.AutoMigrate(&domain.Student{}); err != nil {
		fmt.Println("Student Err")
	} else {
		fmt.Println("Student Suc")
	}

	if err := db.AutoMigrate(&domain.Professor{}); err != nil {
		fmt.Println("Professor Err")
	} else {
		fmt.Println("Professor Suc")
	}

	if err := db.AutoMigrate(&domain.Class{}); err != nil {
		fmt.Println("Class Err")
	} else {
		fmt.Println("Class Suc")
	}
	if err := db.AutoMigrate(&domain.ClassRegistration{}); err != nil {
		fmt.Println("ClassRegistration Err")
	} else {
		fmt.Println("ClassRegistration Suc")
	}

	if err := db.AutoMigrate(&domain.Score{}); err != nil {
		fmt.Println("Score Err")
	} else {
		fmt.Println("Score Suc")
	}

}
