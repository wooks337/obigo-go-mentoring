package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"strconv"
	"todo/Service"
	"todo/domain"
)

var rd *render.Render
var db *gorm.DB

func main() {
	rd = render.New()
	mux := MakeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	log.Println("Started App")

	ConnectDB()

	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}

	err = ConnectDB()
	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
		fmt.Println("=================끝")
	}()

	if err := db.AutoMigrate(&domain.Todo{}); err != nil {
		fmt.Println("Todo Err")
	} else {
		fmt.Println("Todo Suc")
	}
}

func ConnectDB() (err error) {
	dsn := "root:root@tcp(10.28.3.180:3307)/gormMax?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})
	return err
}

func MakeWebHandler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("public")))
	s := r.PathPrefix("/todos").Subrouter()
	s.HandleFunc("", GetTodoListHandler).Methods("GET")
	s.HandleFunc("", PostTodoHandler).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	s.HandleFunc("/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")

	//mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	//mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	//mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	//mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")

	return r
}

func GetTodoListHandler(w http.ResponseWriter, req *http.Request) {

	list, err := Service.GetTodoList(db)

	if err != nil {
		rd.JSON(w, http.StatusInternalServerError, Success{false})
		return
	}
	log.Println(list)
	rd.JSON(w, http.StatusOK, list) //JSON 변환
}

func PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	var todo domain.Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Fatal("err발생 : ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postTodo, err := Service.PostTodo(db, todo)
	if err != nil {
		log.Fatal("err발생 : ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rd.JSON(w, http.StatusCreated, postTodo)
}

func RemoveTodoHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if err := Service.RemoveTodo(db, id); err != nil {
		rd.JSON(w, http.StatusNotFound, Success{false})
	} else {
		rd.JSON(w, http.StatusOK, Success{true})
	}
}

func UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {

	var todo domain.Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Fatal("err발생 : ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	if findTodo := Service.GetTodoById(db, id); findTodo != nil {

		if updateTodo, err := Service.UpdateTodo(db, todo); err != nil {
			rd.JSON(w, http.StatusInternalServerError, Success{false})
		} else {
			fmt.Println(updateTodo)
			rd.JSON(w, http.StatusOK, Success{true})
		}
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}

type Success struct {
	Success bool `json:"success"`
}
