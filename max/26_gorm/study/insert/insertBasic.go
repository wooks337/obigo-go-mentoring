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
		Logger:          logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
		CreateBatchSize: 1000,                                //배치사이즈 설정
	})

	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	db = db.Session(&gorm.Session{CreateBatchSize: 1000}) //배치 사이즈 미리지정

	//일반Insert
	std := Student{Name: "이청용", MajorId: 2}
	result := db.Create(&std)
	fmt.Println("result : ", result.Error)
	fmt.Println(std.Id)

	//배치Insert
	//students := []Student{}
	//students = append(students, Student{Name: "차범근", MajorId: 1, StudentId: 1})
	//students = append(students, Student{Name: "차두리", MajorId: 1, StudentId: 2})
	//result := db.CreateInBatches(&students, 100)
	//fmt.Println("result : ", result.Error)

}

type Student struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
}

func (Student) TableName() string {
	return "student"
}
