package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student // 학생목록 저장하는 맵   (key =Id, value =Student type)
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter() //gorilla/mux 만들기
	// "/students"에 대한 GET 요청 수령 시, GetStudentListHandler() 함수 호출
	//Method()함수로, GET 요청 받을 때만 핸들러 동작
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	//새로운 핸들러 등록//
	//특정 학생 데이터 반환 : GET 메서드로 /students/경로 아래 숫자로 된 경로가 온다면 GetStudentHandler() 함수 호출
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET") //gorilla/mux에 자동으로 id 값 내부 맵에 저장
	//학생데이터 추가
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	//학생데이터 삭제
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student) //임시 데이터 2개 생성
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
	return mux
}

//ID로 정렬하는 인터페이스 구현
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

//GetStudentListHandler() 함수 : 학생정보를 가져와 JSON 포맷으로 변경하는 핸들러
func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) //학생목록을 ID로 정렬
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) //json 포맷으로 변경
}

//학생데이터 조회
//GetStudentHandler() 함수
func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)               //mux.Var()함수 호출하여 인수 가져오고
	id, _ := strconv.Atoi(vars["id"]) //var(["id"]로 id값 가져오기 + 문자열-> 숫자로 변경
	student, ok := students[id]       //student 맵에서 해당 학생 데이터 여부 확인
	if !ok {
		w.WriteHeader(http.StatusNotFound) //id에 해당하는 학생이 없으면 404 에러
		return
	}
	// if ok ~ http status 200
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json") //학생 데이터 json 포맷으로 변경
	json.NewEncoder(w).Encode(student)
}

//학생 데이터 추가
//PostStudentHandler() 함수
func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student) //요청에 포함된 json 데이터를 Student 타입으로 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //변환 관정에서 에러 발생 시 BadRequest 코드 반환
		return
	}
	lastId++ //id 증가시킨 후 맵에 등록
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated) //StatusCreated 코드로 학생 데이터 추가를 알림
}

//학생 데이터 삭제
//DeleteStudentHandler() 함수
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //id에 해당하는 학생 없으면 404 에러 반환
		return
	}
	delete(students, id)         // 학생 맵에서 해당 id 삭제
	w.WriteHeader(http.StatusOK) //statusOK 반환
}

//main 함수만들고 3000에서 입력 대기
func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
