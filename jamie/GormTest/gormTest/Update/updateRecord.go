package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obigo-go-mentoring/jamie/GormTest/gormTest/domain"
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

	//====Update Basic====
	////Save() 메서드
	var prof domain.Prof
	//db.First(&prof)

	//prof.Name = "마담신"
	//prof.Age = 36
	//db.Save(&prof)

	////단일 컬럼 Update
	//db.Model(&domain.Prof{}).Where("Age = ?", 35).Update("name", "이 향")
	//db.Model(&prof).Update("dept_id", 2) //db.First(&prof)를 참조하여 모델의 기본값을 조건으로 사용
	//db.Model(&prof).Where("dept_id = ?", 2).Update("country", "france")

	////복수 컬럼 Update
	////Struct
	//db.Model(&prof).Updates(domain.Prof{Name: "신정아", DeptID: 7})

	////Map
	//db.Last(&prof)
	//db.Model(&prof).Updates(map[string]interface{}{
	//	"name": "기믕운", "country": "algerie",
	//})

	////선택한 필드 Update
	db.Find(&prof).Where("age = ?", 55)
	db.Model(&prof).Select("*").Omit("ProfID", "Name").Update(domain.Prof{
		Age:     48,
		Gender:  M,
		Country: "tunisie",
		DeptID:  8,
	})
}
