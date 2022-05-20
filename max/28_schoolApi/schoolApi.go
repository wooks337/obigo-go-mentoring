package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"school/domain"
	"school/service"
)

var db *gorm.DB
var rd *render.Render

func main() {
	rd = render.New()
	mux := MakeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	log.Println("Started App!")

	var err error
	db, err = ConnectDB()
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)

	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	//DbMigrate(db)

	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}

}

func DbMigrate(db *gorm.DB) {

	if err := db.AutoMigrate(&domain.MajorDepartment{}); err != nil {
		fmt.Println("MajorDepartment Err")
	} else {
		fmt.Println("MajorDepartment Suc")
	}

	if err := db.AutoMigrate(&domain.Major{}); err != nil {
		fmt.Println("Major Err")
	} else {
		fmt.Println("Major Suc")
	}

	if err := db.AutoMigrate(&domain.Student{}); err != nil {
		fmt.Println("Student Err")
	} else {
		fmt.Println("Student Suc")
	}

	if err := db.AutoMigrate(&domain.Professor{}); err != nil {
		fmt.Println("Professor Err")
	} else {
		fmt.Println("Professor Suc")
	}

	if err := db.AutoMigrate(&domain.Class{}); err != nil {
		fmt.Println("Class Err")
	} else {
		fmt.Println("Class Suc")
	}
	if err := db.AutoMigrate(&domain.ClassRegistration{}); err != nil {
		fmt.Println("ClassRegistration Err")
	} else {
		fmt.Println("ClassRegistration Suc")
	}

	if err := db.AutoMigrate(&domain.Score{}); err != nil {
		fmt.Println("Score Err")
	} else {
		fmt.Println("Score Suc")
	}
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(10.28.3.180:3307)/schoolMax?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})
	return db, err
}

func MakeWebHandler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("public")))
	subDepartment := r.PathPrefix("/departments").Subrouter()
	subDepartment.HandleFunc("", GetDepartmentListHandler).Methods("GET")

	subMajor := r.PathPrefix("/majors").Subrouter()
	subMajor.HandleFunc("", GetMajorsListHandler).Methods("GET")
	return r
}

func GetMajorsListHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Majors")
}

func GetDepartmentListHandler(w http.ResponseWriter, req *http.Request) {

	departmentList := service.GetDepartmentList(db)

	rd.JSON(w, http.StatusOK, departmentList)
}
