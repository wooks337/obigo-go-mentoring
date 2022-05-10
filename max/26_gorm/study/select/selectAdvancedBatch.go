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

	films := []_6_gorm.Film{}
	db.Where("film_id <= ?", 100).FindInBatches(&films, 50, func(tx *gorm.DB, batch int) error {
		for _, film := range films {
			fmt.Println(film.Film_id)
		}
		//tx.Save(&films)
		fmt.Println("갯수 : ", tx.RowsAffected)

		//batch // Batch 1, 2, 3 // 뭐지?
		return nil
	})
}
