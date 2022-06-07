package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//func main() {
//	//======localhost======
//	dsn := "root:jamiekim@(localhost:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		panic(err)
//	}

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	type Result struct {
		Rating string
		Total  float64
	}

	//쿼리 결과 담을 슬라이스
	//var film = make([]domain.Film, 1)
	var res = make([]Result, 1)

	////Group By
	////SELECT rating, round(sum(rental_rate)) as total FROM `film` GROUP BY `rating` ORDER BY total
	//db.Model(&domain.Film{}).
	//	Select("rating, round(sum(rental_rate)) as total").
	//	Group("rating").
	//	Order("total").
	//	Find(&res)

	////Group By + Having
	////SELECT rating, round(sum(rental_rate)) as total FROM `film` GROUP BY `rating` HAVING rating = 'R'
	//db.Model(&domain.Film{}).
	//	Select("rating, round(sum(rental_rate)) as total").
	//	Group("rating").
	//	Having("rating = ?", "R").
	//	Find(&res)

	////SELECT rental_rate, round(sum(replacement_cost)) as total FROM `film` GROUP BY `rental_rate`
	//rows, err := db.Table("film").
	//	Select("rental_rate, round(sum(replacement_cost)) as total").
	//	Group("rental_rate").
	//	Rows()
	//defer rows.Close()
	//for rows.Next() {
	//	res2 := Result{}
	//	rows.Scan(&res2.Rating, &res2.Total)
	//	log.Println(res2)
	//}

	//SELECT rating, round(sum(replacement_cost)) as total FROM `film` GROUP BY `rating` HAVING rating LIKE '%G%'
	//db.Table("film").Select("rating, round(sum(replacement_cost)) as total").
	//	Group("rating").
	//	Having("rating LIKE ?", "%G%").
	//	Scan(&res)
	//log.Println(res)
	//db.Model(&domain.Film{}).Select("rating, round(sum(replacement_cost)) as total").
	//	Group("rating").
	//	Having("rating LIKE ?", "%G%").
	//	Scan(&res)
	//log.Println(res)
}
