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

	db.Where("rating = ?", "PG").Find(&films)                           //equal
	db.Where("Rating IN ?", []string{"PG", "R"}).Find(&films)           //IN
	db.Where("Rating Like ?", "%G%").Find(&films)                       //Like
	db.Where("rental_rate = ? AND rating = ?", 2.99, "PG").Find(&films) //AND
	db.Where("rental_rate > ?", 2.99).Find(&films)                      //부등호
	db.Where("film_id Between ? AND ?", 3, 10).Find(&films)             //Between

	for _, film := range films {
		fmt.Println(film.Film_id, " ", film.Title, " ", film.Rating, " ", film.Length, " ", film.Rental_rate)
	}

}
