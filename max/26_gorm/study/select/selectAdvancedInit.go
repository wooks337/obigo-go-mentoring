package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "root:root@tcp(10.28.3.180:3307)/sakila?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})

	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	//=====못 찾았을 때 INIT======
	//해당 컬럼으로 못 찾으면 그 값으로 초기화 됨1
	//film := _6_gorm.Film{}
	//db.FirstOrInit(&film, _6_gorm.Film{Title: "not_found"})
	//fmt.Println(film)

	//해당 컬럼으로 못 찾으면 그 값으로 초기화 됨2
	//film2 := _6_gorm.Film{}
	//db.Where(_6_gorm.Film{Title: "not_found"}).FirstOrInit(&film2)
	//fmt.Println(film2)

	//해당 컬럼으로 못 찾으면 그 값으로 초기화 + 필드값 설정 함
	//film3 := _6_gorm.Film{}
	//db.Attrs(_6_gorm.Film{Length: 120}).FirstOrInit(&film3, _6_gorm.Film{Title: "not_found"})
	//db.Attrs("length", 120).FirstOrInit(&film3, _6_gorm.Film{Title: "not_found"})
	//db.Attrs(map[string]interface{}{
	//	"length":  120,
	//	"film_id": 50,
	//}).FirstOrInit(&film3, _6_gorm.Film{Title: "not_found"})
	//
	//fmt.Println(film3)

	//======찾았을 때 INIT======
	//해당 컬럼 찾아도 그 값으로 설정 됨
	//film4 := _6_gorm.Film{}
	//db.Assign("description", "none").FirstOrInit(&film4, _6_gorm.Film{Title: "ACADEMY DINOSAUR"})
	//db.Assign(_6_gorm.Film{Description: "none", Length: 120}).FirstOrInit(&film4, _6_gorm.Film{Title: "ACADEMY DINOSAUR"})
	//fmt.Println(film4)

	//======못 찾았을 때 Insert======
	//해당 컬럼 못 찾으면 Insert 함
	//film5 := _6_gorm.Film{}
	//db.FirstOrCreate(&film5, _6_gorm.Film{Title: "반지의 제왕", Length: 120, Original_language_id: 1, Language_id: 1})
	//db.Where().FirstOrCreate(&film5)
	//fmt.Println(film5)

	//해당 컬럼 못 찾으면 Insert 하는데 Attrs 값으로 수정됨
	//film6 := _6_gorm.Film{}
	//db.Attrs("description", "none").FirstOrCreate(&film6, _6_gorm.Film{Title: "반지의 제왕3", Length: 120, Original_language_id: 1, Language_id: 1})
	//fmt.Println(film6)

	//해당 컬럼 찾아도 Assign 값으로 수정
	//film7 := _6_gorm.Film{}
	//db.Assign("description", "none1").Where(_6_gorm.Film{Title: "반지의 제왕3", Length: 120, Original_language_id: 1, Language_id: 1}).FirstOrCreate(&film7)
	//fmt.Println(film7)

}
