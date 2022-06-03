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

	////단일 맵 입력
	//db.Model(&domain.Student{}).Create(map[string]interface{}{
	//	"Name": "김정운", "Age": 27, "Gender": "F", "DeptID": 1,
	//})

	////복수 맵 입력
	//db.Model(&domain.Student{}).Create([]map[string]interface{}{
	//	{"Name": "윤이지", "Age": 22, "Gender": "F", "DeptID": 1},
	//	{"Name": "김윤진", "Age": 24, "Gender": "F", "DeptID": 2},
	//	{"Name": "고윤로", "Age": 27, "Gender": "M", "DeptID": 4},
	//	{"Name": "양준서", "Age": 23, "Gender": "M", "DeptID": 5},
	//	{"Name": "김민기", "Age": 23, "Gender": "M", "DeptID": 6},
	//})

}
