package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Bar")
	})

	return mux
}
