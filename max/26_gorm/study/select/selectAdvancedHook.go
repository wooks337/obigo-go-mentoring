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

	films := []FilmHook{}
	db.Where("film_id < ?", 50).Find(&films)
}

func (f *FilmHook) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("===AfterFind===")
	fmt.Println(f)
	//if f.Film_id == 2 {
	//	return errors.New("errr")
	//}
	return
}

type FilmHook struct {
	Film_id              int
	Title                string
	Description          string
	Release_year         int
	Language_id          int
	Original_language_id int
	Rental_duration      int
	Rental_rate          float64
	Length               int
	Replacement_cost     float64
	Rating               string
	Special_features     string
	Last_update          string
}

func (FilmHook) TableName() string {
	return "film"
}
