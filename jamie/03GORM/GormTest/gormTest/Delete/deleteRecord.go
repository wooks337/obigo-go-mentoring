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

	////====Delete Basics=====
	//var stu domain.Student
	//db.First(&stu)
	//db.Delete(&stu)

	//var dept domain.Dept
	//db.First(&dept, 4)
	//db.Where("dept_build = ?", "-").Delete(&dept)

	////PK를 조건으로 Delete
	//db.Delete(&domain.Student{}, 6)
	//db.Delete(&domain.Student{}, "10")
	//db.Delete(&domain.Student{}, []int{14, 15})
}
