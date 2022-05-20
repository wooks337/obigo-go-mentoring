package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
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

	var login []LoginData
	db.Table("login_data").Find(&login)
	log.Println(login, "\n")
}

type LoginData struct { //테스트용 별개 데이터
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	PassWord string
	StuID    string
	DeptID   string
}

func (l *LoginData) AfterFind(tx *gorm.DB) (err error) {
	t := time.Now()
	a := t.Format("2006")
	//fmt.Println(a)

	if l.StuID == "" {
		tx.Statement.SetColumn("stu_id", a)
	}
	return
}
