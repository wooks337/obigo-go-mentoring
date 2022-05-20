package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GORM/GormTest/gormTest/domain"
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

	var stu domain.Student
	db.First(&stu, 26)

	//FOR UPDATE
	db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&stu)

	//FOR SHARE
	db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Find(&stu)

	//FOR UPDATE NOWAIT
	db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	}).Find(&stu)

	//다른 세션에서 접근하는 예제를 어떻게 만들지...
}
