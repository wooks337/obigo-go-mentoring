package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/gormMax?charset=utf8&parseTime=True&loc=Local"
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

	//major := domain.Major{}
	//db.Preload("Students").Find(&major, 1)

	//student (N) : (1) major (N) : (1) major_department
	//student := domain.Student{}
	//db.Preload("Major.MajorDepartment").First(&student)
	//fmt.Println(student.Name)
	//fmt.Println(student.Major.Name)
	//fmt.Println(student.Major.MajorDepartment.Name)

	//student (N) : (1) major (N) : (1) major_department
	//student := domain.Student{}
	//db.Preload(clause.Associations).First(&student)
	//fmt.Println(student.Name)
	//fmt.Println(student.Major.Name)
	//fmt.Println(student.Major.MajorDepartment.Name) //안나옴

	//major call students
	//major2 := domain.Major{}
	//db.Preload("Students").First(&major2)
	//fmt.Println(major2.Name)
	//for _, student := range major2.Students {
	//	fmt.Println(student.Name)
	//}

	//studentJoin := studentWithMaJorWithDepartment{}
	//db.Table("students as s").
	//	Select("s.id as StudentId, s.name as StudentName, m.id as MajorId, m.name as MajorName, md.id as DepartId, md.name as DepartName").
	//	Joins("join majors m ON m.id = s.major_id").
	//	Joins("join major_departments md ON md.id = m.major_department_id").
	//	Where("s.Id = ?", 21).
	//	Scan(&studentJoin)
	//fmt.Println(studentJoin)
}

type studentWithMaJorWithDepartment struct {
	StudentId   int
	StudentName string
	MajorId     int
	MajorName   string
	DepartId    int
	DepartName  string
}
