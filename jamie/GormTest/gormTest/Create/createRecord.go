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

	//====Create Basic====
	//d1 := domain.Dept{DeptName: "프랑스 학과", DeptBuild: "어문관"}
	//d2 := domain.Dept{DeptName: "컴퓨터 공학과", DeptBuild: "공학관"}

	////기본 create
	//db.Create(&d1)
	////특정 필드 지정하여 데이터 저장
	//db.Select("dept_name", "dept_build").Create(&d2)
	////특정 필드 제외하고 데이터 저장
	//db.Omit("dept_build").Create(&domain.Dept{DeptName: "지식 콘텐츠 학과"})

	////Batch Insert
	//depts := []domain.Dept{
	//	{DeptName: "루마니아", DeptBuild: "교양관"},
	//	{DeptName: "러시아", DeptBuild: "어문관"},
	//	{DeptName: "철학", DeptBuild: "인경관"},
	//	{DeptName: "언어인지학", DeptBuild: "인경관"},
	//	{DeptName: "경영정보", DeptBuild: "백년관"},
	//	{DeptName: "경영정보", DeptBuild: "백년관"},
	//}
	//db.CreateInBatches(depts, 100)

}
