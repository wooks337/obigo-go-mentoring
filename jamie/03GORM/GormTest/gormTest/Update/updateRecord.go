package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	//====Update Basic====
	////Save() 메서드
	//var prof domain.Prof
	//db.First(&prof)
	//prof.Name = "마담신"
	//prof.Age = 36
	//db.Save(&prof)

	////단일 컬럼 Update
	//var prof domain.Prof
	//db.First(&prof)
	//db.Model(&domain.Prof{}).Where("Age = ?", 35).Update("name", "이 향")
	//db.Model(&prof).Update("dept_id", 2) //db.First(&prof)를 참조하여 모델의 기본값을 조건으로 사용
	//db.Model(&prof).Where("dept_id = ?", 2).Update("country", "france")

	////복수 컬럼 Update
	//var prof domain.Prof
	////Struct
	//db.First(&prof)
	//db.Model(&prof).Updates(domain.Prof{Name: "신정아", DeptID: 7})
	//
	////Map
	//db.Last(&prof)
	//db.Model(&prof).Updates(map[string]interface{}{
	//	"name": "기믕운", "country": "algerie",
	//})

	////선택한 필드 Update
	var prof domain.Prof
	//db.First(&prof)
	///select로 특정 필드 정하고 map으로 변경
	//db.Model(&prof).Select("name").Updates(map[string]interface{}{
	//	"name": "hello",
	//	"age":  18,
	//})
	///Omit으로 특정 필드 제외하고 map으로 나머지 변경
	//db.Model(&prof).Omit("name").Updates(map[string]interface{}{
	//	"age":     18,
	//	"gender":  "M",
	//	"country": "spain",
	//})
	///select로 특정 필드 정하고 구조체로 변경
	//db.Model(&prof).Select("Name", "DeptID").Updates(domain.Prof{
	//	Name:   "김땡땡",
	//	Age:    28,
	//	DeptID: 2,
	//})
	///전체 필드 변경
	//db.Last(&prof)
	//db.Model(&prof).Select("*").Updates(domain.Prof{
	//	Name:    "김응운",
	//	Age:     55,
	//	Gender:  "M",
	//	Country: "north korea",
	//	DeptID:  8,
	//})
	///특정 필드 제외하고 전체 필드 변경
	//db.First(&prof)
	//db.Model(&prof).Select("*").Omit("Name").Updates(domain.Prof{
	//	Name:    "신정아",
	//	Age:     33,
	//	Gender:  "F",
	//	Country: "Finland",
	//	DeptID:  1,
	//})
}
