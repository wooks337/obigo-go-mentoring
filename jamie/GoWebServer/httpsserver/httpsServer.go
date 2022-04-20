package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil)
	//HTTPS 서버 시작
	if err != nil {
		log.Fatal(err)
	}
}
