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
	//s1 := Student{}
	//db.First(&s1, 33)
	//fmt.Println(s1)
	//s1.MajorId = 2
	//db.Save(&s1)
	//fmt.Println(s1)

	//Select 후 바로Update
	//s2 := Student{}
	//db.Model(&Student{}).Where("id=?", 33).Update("major_id", 1)
	//db.First(&s2, 33)
	//fmt.Println(s2)

	//DB조회한 호출한 데이터로 바로 Update
	//s3 := Student{}
	//db.First(&s3, 33)
	//db.Model(&s3).Update("major_id", 2)
	//fmt.Println(s3)

	//여러 컬럼 Update
	//s4 := Student{}
	//db.First(&s4, 33)
	//db.Model(&s4).Updates(Student{Name: "김용수", MajorId: 1})
	//fmt.Println(s4)

	//여러 컬럼 Update From Map
	//s5 := Student{}
	//db.First(&s5, 33)
	//db.Model(&s5).Updates(map[string]interface{}{
	//	"name":     "최용수",
	//	"major_id": 2,
	//})
	//fmt.Println(s5)

	//객체에서 특정 필드만 Update
	s6 := Student{}
	newStudent := Student{Name: "최용수", MajorId: 1, StudentId: 222}
	db.First(&s6, 33)
	//db.Model(&s6).Select("name", "major_id").Updates(newStudent)	//Select로 특정 필드만 변경
	//db.Model(&s6).Omit("student_id").Updates(newStudent) //Omit으로 특정 필드 제외 하고 변경
	//db.Model(&s6).Select("*").Updates(newStudent) //모든 필드 변경
	db.Model(&s6).Select("*").Omit("id", "student_id").Updates(newStudent) //모든 필드에서 Omit제외 변경

	fmt.Println(s6)

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
