package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type LoginData struct { //테스트용 별개 데이터
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	PassWord string
	StuID    string
	DeptID   string
}

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	//====Update Hooks====
	//var log LoginData
	//db.First(&log)
	//res := db.Model(&log).Update("name", "admin")
	//fmt.Println(res.Error)
	//
	//db.First(&log, 2)
	//db.Model(&log).Update("Name", "김호랭")
	//db.Model(&LoginData{}).Where("name = ?", "김기린").UpdateColumn("StuID", "")

	//========새로운 레코드 주입 시 비밀번호 암호화==========
	var user LoginData
	db.First(&user, 2)
	db.Save(&LoginData{Name: "name1", PassWord: "123"})
}

////Hooks로 에러 출력
//func (l *LoginData) BeforeSave(tx *gorm.DB) (err error) {
//	fmt.Println(l)
//	if l.StuID == "admin" {
//		return errors.New("admin account cannot be updated")
//	}
//	return
//}

//Hooks로 값변환
//func (l *LoginData) BeforeUpdate(tx *gorm.DB) (err error) {
//	t := time.Now()
//	a := t.Format("2006")
//	fmt.Println(a)
//
//	fmt.Println(l)
//	if l.StuID == ":" {
//		tx.Statement.SetColumn("stu_id", a)
//	}
//	return
//}

////데이터 주입시 비밀번호 암호화
func (l *LoginData) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("before save")
	fmt.Println(l.PassWord)
	if l.PassWord != "" {
		hash, err := MakePassword(l.PassWord)
		if err != nil {
			return nil
		}
		tx.Statement.SetColumn("PassWord", hash)
	}
	fmt.Println(l.PassWord)
	return
}

// MakePassword : Encrypt user password
func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
