package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	var students []StudentReturnData
	//모두 조회
	//db.Model(&students).Clauses(clause.Returning{}).Where("id IN ?", []int{27, 28, 29, 30}).Update("age", 21)
	db.Where("id IN ?", []int{27, 28, 29, 30}).Find(&students)
	db.Model(&students).Clauses(clause.Returning{}).Update("student_id", gorm.Expr("id * ? + major_id", 10))

	//특정 컬럼만 조회
	//db.Model(&students).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "name"}, {Name: "age"}}}).
	//	Where("id IN ?", []int{27, 28, 29, 30}).Update("age", 20)

	fmt.Println(students)

}

type StudentReturnData struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
	Age       int
}

func (StudentReturnData) TableName() string {
	return "student"
}
