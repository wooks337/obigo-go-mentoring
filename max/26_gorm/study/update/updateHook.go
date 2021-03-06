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

	//s := StudentHook{}
	//db.Model(&StudentHook{}).Where("id>=?", 29).Update("major_id", 2)
	//db.First(&s, 33)
	//fmt.Println(s)

	students := []StudentHook{}
	db.Where("id>=?", 29).Find(&students)
	//db.Model(&students).Updates(&StudentHook{MajorId: 2})
	db.Model(&students).Updates(map[string]interface{}{"student_id": gorm.Expr("id * ? + major_id", 10)})

}

type StudentHook struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
}

func (StudentHook) TableName() string {
	return "student"
}

func (s *StudentHook) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("======", s)
	//s.StudentId = s.Id*10 + s.MajorId
	//if s.Id == 33 {
	//	tx.Statement.SetColumn("major_id", 1)
	//}
	return nil
}
