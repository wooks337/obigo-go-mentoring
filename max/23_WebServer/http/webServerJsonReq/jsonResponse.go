package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, req *http.Request) {
	student := Student{"kim", 18, 97}
	data, _ := json.Marshal(student)                   //Student -> []byte로 변환
	w.Header().Add("content-type", "application/json") //Json포멧 표시
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}
