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

	////====Delete Batch=====

	//db.Where("name LIKE ?", "기뮨진").Delete(&domain.Student{})
	//db.Delete(&domain.Student{}, "country LIKE ?", "%al%")

	////조건식 없으면 없다고 오류 뜸
	//res := db.Delete(&domain.Student{})
	//log.Println(res.Error)

	////rawSQL 사용해서 삭제
	//db.Exec("Delete FROM students") //전체 레코드 삭제

	////AllowGlobalUpdate 모드 허용해서 삭제
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Student{})  //전체 레코드 삭제
}
