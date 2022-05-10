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

	db.Model(&StudentGlobal{}).Update("major_id", 2) //조건없어서 오류

	//db.Model(&StudentGlobal{}).Where("1=1").Update("age", "age+1")
	//db.Exec("UPDATE student SET major_id = ?", 2)
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&StudentGlobal{}).Update("major_id", "2")

}

type StudentGlobal struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
}

func (StudentGlobal) TableName() string {
	return "student"
}
