package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	type Dept struct {
		DeptID    uint `gorm:"primaryKey"`
		DeptName  string
		DeptBuild string
	}
	type Prof struct {
		ProfID  uint `gorm:"primaryKey"`
		Name    string
		Age     int
		Gender  string
		Country string `gorm:"default:'south korea'"`
		Dept    Dept   `gorm:"foreignkey:DeptID"`
	}
	type Student struct {
		StuID   uint `gorm:"primaryKey"`
		Name    string
		Age     int
		Gender  string
		Country string `gorm:"default:'south korea'"`
		Dept    Dept   `gorm:"foreignkey:DeptID"`
	}
	//테이블 생성
	db.AutoMigrate(&Dept{})
	db.AutoMigrate(&Student{}, &Prof{})
	//외래키는 DataGrip에서 따로 지정함...

}
