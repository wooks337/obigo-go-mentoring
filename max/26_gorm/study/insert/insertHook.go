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

	student := StudentHook{Name: "기성용", MajorId: 2}
	result := db.Create(&student)

	fmt.Println(result)

}

type StudentHook struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
}

func (StudentHook) TableName() string {
	return "student"
}

func (s *StudentHook) AfterCreate(tx *gorm.DB) (err error) {
	s.StudentId = s.Id*10 + s.MajorId
	return nil
}
