package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	http.ListenAndServe(":3000", MakeWebHandler()) //import "net/http"
}

func MakeWebHandler() http.Handler {
	students = make(map[int]Student) //map 생성

	router := mux.NewRouter() //mux의 라우터 생성
	router.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	router.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	router.HandleFunc("/students", PostStudentHandler).Methods("POST")
	router.HandleFunc("/students/{id:[0-9]+}", DeleteHandler).Methods("DELETE")
	router.HandleFunc("/students/{id:[0-9]+}", PutHandler).Methods("PUT")

	return router
}

func GetStudentListHandler(w http.ResponseWriter, req *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)                       //상태 값
	w.Header().Set("Content-Type", "application/json") //header 값 설정
	json.NewEncoder(w).Encode(list)                    //import "encoding/json"
}

func GetStudentHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) //Vars를 이용하여 Path변수 확인
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //id값 없는 경우 NotFount 반환
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, req *http.Request) {
	var student Student
	err := json.NewDecoder(req.Body).Decode(&student) //Json데이터 Decode
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Decode 실패시 BadRequest 반환
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated) //생성완료 Created 반환
}

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //id값 없는 경우 NotFount 반환
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK)
}

func PutHandler(w http.ResponseWriter, req *http.Request) {
	var updateStudent Student
	err := json.NewDecoder(req.Body).Decode(&updateStudent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Decode 실패시 BadRequest 반환
		return
	}

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	oldStudent, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //id값 없는 경우 NotFount 반환
		return
	} else {
		oldStudent.Name = updateStudent.Name
		oldStudent.Age = updateStudent.Age
		oldStudent.Score = updateStudent.Score
		students[oldStudent.Id] = oldStudent //Update 후 저장
		w.WriteHeader(http.StatusOK)         //완료 OK 반환
	}
}

type Student struct {
	Id    int    `json:"id"` //결과 출력시 보여지는 이름
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

var students map[int]Student //map선언 key:int, value:student
var lastId int               //마지막 인덱스 저장

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}
