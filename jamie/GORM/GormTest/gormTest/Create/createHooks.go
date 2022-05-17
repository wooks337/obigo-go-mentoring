package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//======TeamServer======
type LoginData struct { //테스트용 별개 데이터
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	PassWord string
	StuID    string
	DeptID   string
}

func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&LoginData{})
	db.Create(&LoginData{Name: "관리자"})
	db.Create(&LoginData{Name: "김기린", StuID: "2015", DeptID: "0035"})
}

//AfterCreate Id = 1, Name = Admin, Pw = 1234, else, Pw = stuID+deptID
func (l *LoginData) AfterCreate(tx *gorm.DB) (err error) {
	if l.Id == 1 {
		tx.Model(l).Update("Name", "Admin").Update("PassWord", 1234)
	} else {
		tx.Model(l).Update("PassWord", l.StuID+l.DeptID)
	}
	return
}
