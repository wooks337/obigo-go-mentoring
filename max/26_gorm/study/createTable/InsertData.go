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

	//학부 추가
	//departments := []domain.MajorDepartment{}
	//departments = append(departments, domain.MajorDepartment{Name: "IT학부"})
	//departments = append(departments, domain.MajorDepartment{Name: "언어학부"})
	//db.Create(&departments)

	//학과 추가
	//departments2 := []domain.MajorDepartment{}
	//db.Find(&departments2)
	//majors := []domain.Major{}
	//majors = append(majors, domain.Major{Name: "컴퓨터공학과", MajorDepartment: departments2[0]})
	//majors = append(majors, domain.Major{Name: "정보통신학과", MajorDepartment: departments2[0]})
	//majors = append(majors, domain.Major{Name: "영어학과", MajorDepartment: departments2[1]})
	//majors = append(majors, domain.Major{Name: "중국어학과", MajorDepartment: departments2[1]})
	//db.Create(&majors)

	//학생 추가
	//majors2 := []domain.Major{}
	//db.Find(&majors2)
	//students := []domain.Student{}
	//students = append(students, domain.Student{Name: "김승규", Age: 24, Major: majors2[0]})
	//students = append(students, domain.Student{Name: "조현우", Age: 23, Major: majors2[0]})
	//students = append(students, domain.Student{Name: "구성윤", Age: 24, Major: majors2[1]})
	//students = append(students, domain.Student{Name: "김영권", Age: 24, Major: majors2[1]})
	//students = append(students, domain.Student{Name: "김민재", Age: 24, Major: majors2[2]})
	//students = append(students, domain.Student{Name: "정승현", Age: 25, Major: majors2[2]})
	//students = append(students, domain.Student{Name: "박지수", Age: 24, Major: majors2[3]})
	//students = append(students, domain.Student{Name: "권경원", Age: 24, Major: majors2[3]})
	//students = append(students, domain.Student{Name: "홍철", Age: 24, Major: majors2[0]})
	//db.Create(&students)

	//major, student 동시 추가 Associations적용
	//var newMajor domain.Major
	//newStudents := []domain.Student{}
	//newStudents = append(newStudents, domain.Student{Name: "홍명보", Age: 22, Major: newMajor})
	//newStudents = append(newStudents, domain.Student{Name: "황선홍", Age: 23, Major: newMajor})
	//newMajor = domain.Major{Name: "일본어학과", Students: newStudents, MajorDepartmentId: 2}
	//db.Create(&newMajor)

	//major, student 동시 추가 Associations미적용
	//var newMajor domain.Major
	//newStudents := []domain.Student{}
	//newStudents = append(newStudents, domain.Student{Name: "홍금보", Age: 22, Major: newMajor})
	//newStudents = append(newStudents, domain.Student{Name: "황적홍", Age: 23, Major: newMajor})
	//newMajor = domain.Major{Name: "독일어학과", Students: newStudents, MajorDepartmentId: 2}
	//db.Omit("Students").Create(&newMajor)

}
