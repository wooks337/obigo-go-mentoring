package main

import "net/http"

func main() {
	// "/"경로에 대한 요청이 올때, static폴더 아래에 있는 파일 제공하는 파일 서버
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", nil)
}
