package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.NewServeMux() 이용하여 새로운 ServeMux 인스턴스 생성
	mux := http.NewServeMux()
	//HandleFunc() 메소드 이용하여 핸들러 인스턴스에 핸들러 등록
	//
	//http.HandleFunc : DefaultServeMux에 핸들러 등록
	//mux.HandleFunc  : ServeMux 인스턴스에 핸들러 등록
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	//ListenAndServe()함수 호출시, ServeMux 인스턴스 넣어 새로운 인스턴스 사용
	http.ListenAndServe(":3000", mux)
}
