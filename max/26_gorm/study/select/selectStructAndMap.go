package main

import (
	"fmt"
	stru "goMod"
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

	films := make([]stru.Film, 1)

	//구조체
	db.Where(&stru.Film{Rating: "PG", Rental_rate: 2.99}).Find(&films)
	//default값 인식못함
	db.Where(&stru.Film{Rating: "PG", Rental_rate: 0}).Find(&films)
	//구조체에서 특정필드로만 검색하기
	db.Where(&stru.Film{Rating: "PG", Rental_rate: 2.99}, "rating").Find(&films)

	//맵으로 검색
	//db.Where(map[string]interface{}{"rating": "PG", "rental_rate": 4.99}).Find(&films)

	for _, film := range films {
		fmt.Println(film.Film_id, " ", film.Title, " ", film.Rating, " ", film.Length, " ", film.Rental_rate)
	}

}
