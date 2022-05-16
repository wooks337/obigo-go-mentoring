package main

import (
	"fmt"
	m "github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"obigo-go-mentoring/jamie/GormRestAPI/Controller"
	"obigo-go-mentoring/jamie/GormRestAPI/database"
)

//=====MySQL 서버 연결=====//
var DB *gorm.DB
var err error
var dsn = "root:root@(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}
	log.Println("Connected to Database...")
	//=====DB 생성=====//
	DB.AutoMigrate(&database.Student{})
	//s1 := database.Student{Name: "김정운", Subject: "국어", Score: 55}
	//s2 := database.Student{Name: "김주연", Subject: "영어", Score: 95}
	//DB.Create(&s1)
	//DB.Create(&s2)
}

func main() {
	Connect()
	//=====RECORDS 삽입=====//
	//db.Model(&database.Student{}).Create([]map[string]interface{}{
	//	{"StudentID": 1, "Name": "김정운", "Score": 55, "SubjectID": 1},
	//	{"StudentID": 2, "Name": "박건형", "Score": 95, "SubjectID": 3},
	//	{"StudentID": 3, "Name": "서경수", "Score": 75, "SubjectID": 2},
	//	{"StudentID": 4, "Name": "이아진", "Score": 85, "SubjectID": 2},
	//})
	//db.Model(&database.Subject{}).Create([]map[string]interface{}{
	//	{"SubName": "국어", "StudentID": 1},
	//	{"SubName": "영어", "StudentID": 1},
	//	{"SubName": "수학", "StudentID": 2},
	//})
	//localHost:3000에서 대기

	http.ListenAndServe(":3000", MakeWebHandler())

}

//라우터 등록
func MakeWebHandler() http.Handler {
	r := m.NewRouter()
	r.HandleFunc("/students", Controller.GetStudents).Methods("GET")

	return r
}
