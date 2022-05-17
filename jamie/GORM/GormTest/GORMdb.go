package _gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//=========MySQL 연결==========
func main() {
	dsn := "root:jamiekim@(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	/*
		func main() {
		//	dsn := "root:root@(10.28.3.180:3307)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
		//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//		Logger: logger.Default.LogMode(logger.Info),
		//	})
		//	if err != nil {
		//		panic(err)
		//	}
	*/
	//========테이블 선언========
	type UserInfo struct {
		ID     uint
		Name   string
		Gender string `gorm:"default:'female'"`
		Hobby  string
	}
	////========테이블 생성========
	//db.AutoMigrate(&UserInfo{})

	//========CREATE========
	user := UserInfo{
		Name:   "James",
		Gender: "male",
		Hobby:  "basketball",
	}
	db.Create(&user)

	db.Select("Name", "Gender", "Hobby").Create(&UserInfo{Name: "Lisa", Gender: "female", Hobby: "football"})

	db.Omit("Name").Create(&UserInfo{Hobby: "skiing"})

	db.Create(&UserInfo{Name: "Martine", Gender: "male", Hobby: "travel"})
	db.Create(&UserInfo{Name: "Jamie", Gender: "female", Hobby: "running"})
	db.Create(&UserInfo{Name: "Randy", Gender: "male", Hobby: "drawing"})
	//========일괄 삽입(batch Insert)========
	var users = []UserInfo{{Name: "Lilly"}, {Name: "Severus"}, {Name: "Limus"}}
	db.Create(&users)

	////========READ========
	//user := &UserInfo{}
	//users := []UserInfo{}
	//
	//db.Take(&user)
	//fmt.Printf("take: %#v\n", user)
	//
	//db.First(&user)
	//fmt.Printf("first: %#v\n", user)
	//
	//db.First(&user, 5)
	//fmt.Println("first:", user)
	//
	//db.Last(&user)
	//fmt.Printf("last: %#v\n", user)
	//
	//db.Find(&users)
	//fmt.Printf("find: %#v\n", users)

	//String 조건문
	//db.Where("name IN ?", []string{"lilly", "jamie"}).Find(&users)
	//fmt.Println("IN:", users)

	//Struct/Map 조건문
	//db.Where(&UserInfo{Name: "Jamie"}, "hobby").Find(&users)
	//fmt.Println(users)

	//인라인
	//db.Find(&users, map[string]interface{}{"gender": "male"})
	//fmt.Println(users)

	////////데이터 수정//////
	//db.Model(&user1).Update("hobby", "crossfit")
	//
	////Delete
	//db.Delete(&user1, 1)
}
