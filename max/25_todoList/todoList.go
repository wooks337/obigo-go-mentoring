package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render

func main() {
	rd = render.New()
	mux := MakeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
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
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	log.Println(list)
	rd.JSON(w, http.StatusOK, list) //JSON 변환
}

func PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	var todo Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Fatal("err발생 : ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

func RemoveTodoHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}

func UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {

	var todo Todo

	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Fatal("err발생 : ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if oldTodo, ok := todoMap[id]; ok {
		oldTodo.Name = todo.Name
		oldTodo.Completed = todo.Completed
		todoMap[id] = oldTodo
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}

type Success struct {
	Success bool `json:"success"`
}

type Todo struct {
	ID        int    `json:"id"` //json포맷으로 변환 옵션, Json변환시 ID가 아닌 id로 변환, 생략가능
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todoMap map[int]Todo
var lastID int = 0

type Todos []Todo //정렬위한
func (t Todos) Len() int {
	return len(t)
}
func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}
