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

	//====Update Batch====
	//Model()에서 지정 레코드에 기본 키 값이 없으면 gorm은 일괄 수정을 진행

	////Update w/ Struct
	//db.Model(domain.Student{}).Where("Name = ?", "김정운").Updates(domain.Student{Name: "Jamie", Age: 25})
	////Update w/ Map
	//db.Table("students").Where("stu_id IN ?", []int{3, 4}).Updates(map[string]interface{}{
	//	"Name": "고로", "Age": 18,
	//})
	////AllowGlobalUpdate 모드 --> 조건식 없이 전체 컬럼에 대해 변경
	//res := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&domain.Student{}).Update("country", "algerie")
	//log.Println(res.RowsAffected)

}
