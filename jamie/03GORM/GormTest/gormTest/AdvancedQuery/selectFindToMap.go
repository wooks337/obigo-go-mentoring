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

	//result := map[string]interface{}{}
	//db.Model(&domain.Student{}).First(&result, "stu_id = ?", 25)
	//log.Println(result)

	//var results []map[string]interface{}
	//db.Table("students").Find(&results)
	//log.Println(results)
}
