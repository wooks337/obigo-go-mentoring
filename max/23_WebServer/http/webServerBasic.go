package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello World") //웹 핸들러 등록
		fmt.Fprintln(w, req)
	})

	http.HandleFunc("/query", queryHandler)

	http.ListenAndServe(":3000", nil) //웹 서버 시작
}

func queryHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(query.Get("id"))
	fmt.Fprintf(w, "Hello %s! id : %d", name, id)
}
