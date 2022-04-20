package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Student 구조체
type Student struct {
	Name  string
	Age   int
	Score int
}

//핸들러 인스턴스 생성 함수
func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   //student 객체를 []byte로 변환
	w.Header().Add("content-type", "application/json") //json 포맷 표시
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) //[]byte -> string, 결과 전송
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
