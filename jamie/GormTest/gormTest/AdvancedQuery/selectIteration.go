package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"obigo-go-mentoring/jamie/GormTest/gormTest/domain"
)

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	rows, err := db.Model(&domain.Student{}).Where("age = ?", 23).Rows()
	defer rows.Close() //프로그램 종료전 rows scan 종료

	for rows.Next() {
		var stu domain.Student
		db.ScanRows(rows, &stu)

		log.Println(stu)
	}

}
