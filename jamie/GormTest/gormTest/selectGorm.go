package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Actor struct {
	Actor_id    int
	First_name  string
	Last_name   string
	Last_update string
}

//type Film_actor struct {
//	Actor_id    int
//	Film_id     int
//	Last_update string
//}

//뒤에 s 붙는 문제 해결
func (Actor) TableName() string {
	return "actor"
}

func main() {
	//======localhost======
	dsn := "root:jamiekim@(localhost:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	//======TeamServer======
	//func main() {
	//	dsn := "root:root@(10.28.3.180:3307)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//		Logger: logger.Default.LogMode(logger.Info),
	//	})
	//	if err != nil {
	//		panic(err)
	//	}

	//var First_name Actor
	var actor Actor

	//Take()
	db.Take(&actor)
	fmt.Println(actor)
	//First()
	db.First(&actor)
	fmt.Println(actor)
	db.First(&actor, 10)
	fmt.Println(actor)
	db.First(&actor, "Last_name = ?", "GUINESS")
	fmt.Println(actor)
}
