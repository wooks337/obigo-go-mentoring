package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/max?charset=utf8&parseTime=True&loc=Local"
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

	s1 := StudentSubQuery{}
	db.First(&s1, 33)

	db.Model(&s1).Update("age", db.Model(&MajorDepartmentSubQuery{}).Select("id").Limit(1))

}

type StudentSubQuery struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
	Age       int
}

func (StudentSubQuery) TableName() string {
	return "student"
}

type MajorDepartmentSubQuery struct {
	Id   int    `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"not null"`
}

func (MajorDepartmentSubQuery) TableName() string {
	return "major_department"
}
