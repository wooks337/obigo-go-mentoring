package main

import (
	"encoding/json"
	"fmt"
	m "github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"obigo-go-mentoring/jamie/GORM/GormRestAPI/database"
	"strconv"
)

//=====MySQL 서버 연결=====//
var db *gorm.DB
var err error

func main() {
	var dsn = "root:root@(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"
	//var dsn = "root:jamiekim@(localhost:3306)/jamie?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}
	log.Println("Connected to Database...")

	defer func() {
		base, _ := db.DB()
		base.Close()
		log.Println("Database Closed...")
	}()

	//=====DB 생성=====//
	db.AutoMigrate(&database.Student{})
	//s1 := database.Student{Name: "김정운", Subject: "국어", Score: 55}
	//s2 := database.Student{Name: "김주연", Subject: "영어", Score: 95}
	//DB.Create(&s1)
	//DB.Create(&s2)

	//localHost:3000에서 대기
	http.ListenAndServe(":3000", MakeWebHandler())
}

//라우터 등록
func MakeWebHandler() http.Handler {
	r := m.NewRouter()
	//create
	r.HandleFunc("/students", CreateStudent).Methods("POST")
	//select all
	r.HandleFunc("/students", GetStudents).Methods("GET")
	//select
	r.HandleFunc("/students/{students_id}", GetStudent).Methods("GET")
	//update
	//r.HandleFunc("/students", PostStudent).Methods("POST")
	//delete
	r.HandleFunc("/students/{student_id}", DeleteStudent).Methods("DELETE")

	return r
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var stud database.Student
	err := json.NewDecoder(r.Body).Decode(&stud) //요청에 포함된 json 데이터를 Student 타입으로 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //변환 관정에서 에러 발생 시 BadRequest 코드 반환
		return
	}
	db.Create(&stud)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //StatusCreated 코드로 학생 데이터 추가를 알림
	json.NewEncoder(w).Encode(stud)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var studs []database.Student
	db.Table("students").Find(&studs)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studs) //json 포맷으로 변경
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var studs []database.Student
	db.First(&studs)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json") //학생 데이터 json 포맷으로 변경
	json.NewEncoder(w).Encode(studs)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	inputStudID := params["student_id"]
	id, _ := strconv.ParseUint(inputStudID, 10, 64)
	deleteId := uint(id)

	db.Where("student_id = ?", deleteId).Delete(&database.Student{})
	w.WriteHeader(http.StatusNoContent)
}
