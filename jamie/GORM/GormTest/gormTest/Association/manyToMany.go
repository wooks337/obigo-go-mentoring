package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

//// User has and belongs to many languages, use `user_languages` as join table
//type Language struct {
//	gorm.Model
//	Name string
//}
//
//type User struct {
//	gorm.Model
//	Languages []Language `gorm:"many2many:user_languages;"`
//}

type User struct {
	gorm.Model
	UserName  string
	Languages []*Language `gorm:"many2many:user_languages;"`
}
type Language struct {
	gorm.Model
	LangName string
	Users    []*User `gorm:"many2many:user_languages;"`
}

//======TeamServer======
func main() {
	dsn := "root:root@(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	//테이블 생성
	//db.AutoMigrate(&User{}, &Language{})

	//언어 종류 생성하면서 동시에 유저 생성 - 영어를 사용하는 Jamie
	//language := Language{LangName: "영어"}
	//user := User{Languages: []*Language{&language}, UserName: "Jamie"}
	//db.Create(&user)

	//이때, 상기 코드로 이미 '영어' 가 생성되었기 때문에 아래와 같이 같은 코드로 돌리면 영어가 중복으로 하나 더 생긴다
	//language2 := Language{Name: "영어"}
	//user2 := User{Languages: []*Language{&language2}, UserName: "Max"}
	//
	//db.Create(&user2)

	//위와 같은 상황을 방지하기 위해 한번 생성된 언어 Row를 쿼리로 호출하여 User만 create 한다.
	// - 영어를 사용하는 Max
	//english := Language{}
	//db.Where("id = ?", "1").Find(&english)
	//user2 := User{Languages: []*Language{&english}, UserName: "Max"}
	//db.Create(&user2)

	user := User{}
	db.Preload("Languages").Where("id = ?", 1).Find(&user)
	//db.Joins("user_languages").Where("users.id = ?", 1).Find(&user)
	log.Println(user)
}
