package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// User has one CreditCard, CreditCardID is the foreign key
	type CreditCard struct {
		gorm.Model
		Number string
		UserID uint
	}
	type User struct {
		gorm.Model
		CreditCard CreditCard
	}

	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&CreditCard{})

	var user User
	var card CreditCard
	db.Model(&user).related()
}
