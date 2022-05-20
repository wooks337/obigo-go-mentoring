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

	//====Update SQLExpr====
	//var stu domain.Student
	//db.Last(&stu)
	//db.Model(&stu).Update("Age", gorm.Expr("Age + ?", 5))

	//db.Model(&stu).Updates(map[string]interface{}{
	//	"age": gorm.Expr("age + dept_id * ?", 2),
	//})
}
