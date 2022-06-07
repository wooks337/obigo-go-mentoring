package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/03GORM/GormTest/gormTest/domain"
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

	//====Update Return Changed Data====
	var studs []domain.Student
	db.Model(&studs).Clauses(clause.Returning{}).Where("dept_id = ?", 5).Update("Age", 4)
	//db.Model(&studs).Clauses(clause.Returning{Columns: []clause.Column{{Name: "Name"}, {Name: "Age"}}}).Where("dept_id = ?", 5).Update("Age", gorm.Expr("age *?", 2))
}
