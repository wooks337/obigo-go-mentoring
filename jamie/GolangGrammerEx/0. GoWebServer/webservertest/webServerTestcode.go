package main

import (
	"fmt"
	"net/http"
)

//핸들러 인스턴스를 생성하는 함수 MakeWebHandler 만들기
func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bar")
	})
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
