package main

import "net/http"

func main() {
	// "/"경로에 대한 요청이 올때, static 폴더 아래에 있는 파일 제공하는 파일 서버
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}

//웹브라우저에 링크 입력 시 이미지 출력 : http://localhost:3000/static/1438678565374.jpg
