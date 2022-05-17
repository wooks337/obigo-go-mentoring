package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"obigo-go-mentoring/jamie/GORM/GormTest/gormTest/domain"
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

	var stu domain.Student
	db.Where("name = @name", sql.Named("name", "이다영")).Find(&stu)
	log.Println(stu)
}
