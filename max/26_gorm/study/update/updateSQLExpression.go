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

	//기본
	//db.Model(&StudentSQL{}).Where("1=1").Update("age", gorm.Expr("age + ?", 1))

	//Map
	//db.Model(&StudentSQL{}).Where("major_id = ?", 2).Updates(map[string]interface{}{
	//	"major_id": 3,
	//	"age":      gorm.Expr("age + ?", 1),
	//})

	//UpdateColmn
	//db.Model(&StudentSQL{}).Where("major_id = 3").UpdateColumn("major_id", gorm.Expr("major_id - ?", 1))
}

type StudentSQL struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
	Age       int
}

func (StudentSQL) TableName() string {
	return "student"
}
