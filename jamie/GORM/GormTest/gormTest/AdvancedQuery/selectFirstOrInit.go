package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	////FirstOrInit
	////조회시 첫번째 레코드 반환, 해당 데이터 없을 경우 제시 조건 그대로 반환
	//var stu domain.Student
	//db.FirstOrInit(&stu, domain.Student{Name: "없음"})
	//db.Where(domain.Student{Name: "없음"}).FirstOrInit(&stu)
	//db.FirstOrInit(&stu, map[string]interface{}{"name": "없음"})
	//log.Println(stu)

	////Attrs
	////조회시 일치 레코드를 찾으면 Attributes 무시, 없으면 attr 포함해서 그대로 반환
	//var stu domain.Student
	//db.Where(domain.Student{Name: "없음"}).Attrs(domain.Student{Age: 1}).FirstOrInit(&stu)
	//db.Where(domain.Student{Name: "이다영"}).Attrs(domain.Student{Age: 25}).FirstOrInit(&stu)
	//log.Println(stu)

	////Assign
	////조회시 일치 레코드를 찾으면 Assign 조건대로 결과 반환(실제 DB에서 변경되지 않음), 없으면 Assign 포함해서 그대로 반환
	//var stu domain.Student
	//db.Where(domain.Student{Name: "없음"}).Assign(domain.Student{Age: 1}).FirstOrInit(&stu)
	//db.Where(domain.Student{Name: "이다영"}).Assign(domain.Student{Age: 1}).FirstOrInit(&stu)
	//log.Println(stu)
}
