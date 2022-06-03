package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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

	var results []domain.Student
	db.Table("students").Find(&results)
	//log.Println("==1", results)
	result := db.Where("age = ?", 23).
		FindInBatches(&results, 4, func(tx *gorm.DB, batch int) error {
			for _, result := range results {
				log.Println("==2", result)
			}
			//tx.Save(&results)
			log.Println("==3", tx.RowsAffected)
			log.Println("==4", batch)
			return nil
		})
	log.Println(result.Error)
	log.Println(result.RowsAffected)
}
