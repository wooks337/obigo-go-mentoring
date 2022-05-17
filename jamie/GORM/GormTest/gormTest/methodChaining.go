package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GORM/GormTest/gormTest/domain"
)

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	//이때, db는 새로 초기화된 *gorm.DB -- 재사용 가능
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	var stu domain.Student
	var stu2 domain.Student

	//Chain method, Finisher Method 뒤에 gorm은 초기화된 인수를 제공하는데
	//이것을 그냥 재사용하면 쿼리 결과가 잘못될 수 있으므로 NewSessionMethod로 재사용가능한 쿼리문을 짠다

	//===Polluted instance===//
	qdb := db.Where("name = ?", "없음")
	qdb.Where("age = ?", 23).Find(&stu)
	qdb.Where("age = ?", 33).Find(&stu2)

	//===NewSessionMethod===//
	qdb2 := db.Where("name = ?", "없음").Session(&gorm.Session{})
	//qdb2 := db.Where("name = ?", "없음").WithContext(context.Background())
	//qdb2 := db.Where("name = ?", "없음").Debug()

	qdb2.Where("age = ?", 23).Find(&stu)
	qdb2.Where("age = ?", 33).Find(&stu2)
}
