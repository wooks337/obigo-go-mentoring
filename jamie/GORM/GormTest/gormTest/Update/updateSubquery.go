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

	//====Update SubQuery====
	//var stu domain.Student
	//db.First(&stu)

	//db.Model(&stu).Update("Age", db.Model(&domain.Prof{}).Select("Age").Where("dept_id = ?", 2))
	//db.Table("students as stu").Where("age = ?", 22).Update("name", db.Table("profs as p").Select("name").Where("p.dept_id = stu.dept_id"))
	//db.Table("students as stu").Where("age = ?", 35).Updates(map[string]interface{}{
	//	"name": db.Table("profs as p").Select("name").Where("country = ?", "spain"),
	//})
}
