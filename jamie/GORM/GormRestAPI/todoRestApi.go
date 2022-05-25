package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"obigo-go-mentoring/jamie/GORM/GormRestAPI/database"
)

//전역 변수 선언
var r *render.Render
var todoMap map[int]database.Todo
var lastID int = 0

func main() {
	r = render.New()
	m := MakeWebHandler()  //아래에 내가 만든 핸들러
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(m)        //UseHandler 메서드로 내가 만든 핸들러 감싸서 http 요청 처리 전에 사용자 핸들러 호출
	log.Println("Started App")

	//db 연결
	err := ConnDB()
	if err != nil {
		panic(err)
	}

	//todo db 만들기
	db.AutoMigrate(&database.Todo{})

	//3000번 포트에서 요청 대기
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}

}

//db 연결 함수
func ConnDB() (err error) {
	var dsn = "root:root@(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"
	//var dsn = "root:jamiekim@(localhost:3306)/jamie?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return err
}

//핸들러 연결
func MakeWebHandler() http.Handler {

	todoMap = make(map[int]database.Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))                     //
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")              //전체 목록 반환
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")                //항목 추가
	mux.HandleFunc("/todos/{id:[0-9]+", RemoveTodoHandler).Methods("DELETE") //항목 삭제
	mux.HandleFunc("/todos/{id:[0-9]+", UpdateTodoHandler).Methods("PUT")    //항목 수정
	return mux
}
