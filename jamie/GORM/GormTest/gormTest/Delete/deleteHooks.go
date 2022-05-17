package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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

	////====Delete Hooks=====
	var login LoginData
	db.First(&login)
	res := db.Delete(&login)
	log.Println(res.Error)
}

func (l *LoginData) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println(l)
	if l.StuID == "admin" {
		return errors.New("admin user not allowed to delete")
	}
	return
}

type LoginData struct { //테스트용 별개 데이터
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	PassWord string
	StuID    string
	DeptID   string
}
