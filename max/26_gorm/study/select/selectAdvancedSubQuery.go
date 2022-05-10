package main

import (
	"fmt"
	_6_gorm "goMod"
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

	//서브쿼리 기본
	//films := []_6_gorm.Film{}
	//db.Where("length > (?)", db.Table("film").Select("AVG(length)")).Find(&films)
	//fmt.Println(films)

	//서브쿼리 변수로 넣기
	//films2 := []_6_gorm.Film{}
	//subQuery := db.Table("film").Select("AVG(length)")
	//db.Where("length > (?)", subQuery).Find(&films2)
	//fmt.Println(films2)

	//From절 서브쿼리 (인라인뷰)
	//films3 := []_6_gorm.Film{}
	//db.Table("(?) as f", db.Model(&_6_gorm.Film{}).Select("film_id", "title", "length")).
	//	Where("f.length > ?", 120).Find(&films3)
	//fmt.Println(films3)

	//From절 서브쿼리 변수로 넣기
	films4 := []_6_gorm.Film{}
	sub1 := db.Model(&_6_gorm.Film{}).Select("film_id", "title").Where("film_id Between ? AND ?", 1, 10)
	db.Table("(?) as s1", sub1).Find(&films4)
	fmt.Println(films4)
}
