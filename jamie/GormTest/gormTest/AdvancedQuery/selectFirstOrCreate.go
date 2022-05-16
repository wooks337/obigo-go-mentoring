package main

import (
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

	////FirstOrCreate
	////조회 시 첫번째 레코드를 반환, 해당 데이터 없을 경우 제시 조건 새로운 레코드 생성
	//var login LoginData
	//result := db.FirstOrCreate(&login, LoginData{Name: "없음"})
	//result := db.Where(LoginData{Name: "오덕구"}).FirstOrCreate(&login)
	//log.Println(login)
	//log.Println(result.RowsAffected)
	//
	////Attrs
	////조회시 일치 레코드를 찾으면 Attributes 무시, 없으면 attr 포함해서 새로운 레코드 생성
	//var login LoginData
	//db.Where(LoginData{Name: "non"}).Attrs(LoginData{StuID: "2022"}).FirstOrCreate(&login)
	//db.Where(LoginData{Name: "non"}).Attrs(LoginData{StuID: "1111"}).FirstOrCreate(&login)
	//log.Println(login)
	//
	////Assign
	////조회시 일치 레코드를 찾으면 Assign 조건대로 결과 반환(실제 DB에서 변경됨), 없으면 Assign 포함해서 새로운 레코드 생성
	//var login LoginData
	//db.Where(LoginData{Name: "None"}).Assign(LoginData{StuID: "2023"}).FirstOrCreate(&login)
	//db.Where(LoginData{Name: "None"}).Assign(LoginData{StuID: "1111"}).FirstOrCreate(&login)
	//log.Println(login)
}
