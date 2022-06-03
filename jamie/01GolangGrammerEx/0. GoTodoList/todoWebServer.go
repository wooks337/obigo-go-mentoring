package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

//할일 정보 담는 구조체
type Todo struct {
	ID        int    `json:"id,omitempty"` //json 포맷 변환 옵션(항목 이름 : id, 생략 가능)
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

type Success struct {
	Success bool `json:"success"`
}

//전역 변수 선언
var rd *render.Render
var todoMap map[int]Todo
var lastID int = 0

func main() {
	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러 : 터미널에 로그 표시, public 폴더 파일 서버 자동 동작
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

//웹 서버 핸들러 생성
func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")              //전체 목록 반환
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")                //항목 추가
	mux.HandleFunc("/todos/{id:[0-9]+", RemoveTodoHandler).Methods("DELETE") //항목 삭제
	mux.HandleFunc("/todos/{id:[0-9]+", UpdateTodoHandler).Methods("PUT")    //항목 수정
	return mux
}

//항목 추가 함수
func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo) //JSON 데이터 -> Todo 객체
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++ //새로운 ID 등록
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

//항목 삭제 함수
func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}

//항목 수정 함수
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusBadRequest, Success{false})
	}
}

//ID로 정렬하는 인터페이스
type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}
func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}
func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)                 //정렬 함수 Sort()
	rd.JSON(w, http.StatusOK, list) //ID로 정렬하여 render 패키지로 전체 목록 JSON 포맷 반환
}
