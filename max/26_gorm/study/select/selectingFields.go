package main

import (
	"fmt"
	stru "goMod"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/sakila?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	films := make([]stru.Film, 1)

	//필드명
	db.Select("film_id", "title").Where("rental_rate > ?", 2.99).Find(&films) //부등호

	//구조체
	db.Select([]string{"film_id", "title"}).Where("rental_rate > ?", 2.99).Find(&films) //부등호

	//Coalesce 사용
	db.Table("film").Select("COALESCE(rating, ?)", -1).Rows()

	for _, film := range films {
		fmt.Println(film.Film_id, " ", film.Title, " ", film.Rating, " ", film.Length, " ", film.Rental_rate)
	}

}
