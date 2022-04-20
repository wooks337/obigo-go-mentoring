package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//http.Request의 URL의 Query() 메소드로 쿼리 인수를 가져온다
func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    //쿼리 인수 가져오기
	name := values.Get("name") //특정 키 값이 있는지 확인
	if name == "" {
		name = "World"
	}
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	//id 키의 쿼리값을 가져와서 strconv.Atoi()함수로 int 타입 변환
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler) //"/bar" 경로에 핸들러 함수 등록
	http.ListenAndServe(":3000", nil)
}

//http://localhost:3000/bar HTTP요청 수신 시 barHandler() 함수 호출
