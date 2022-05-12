package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GormTest/gormTest/domain"
)

//======TeamServer======
type APIStu struct {
	Name   string
	DeptID uint
}

func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	var s = make([]APIStu, 1)
	db.Model(&domain.Student{}).Limit(5).Find(&s)

	for _, s := range s {
		fmt.Println(s.Name, " ", s.DeptID)
	}

}
