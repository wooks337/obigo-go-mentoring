package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//인덱스 경로 "/" 경로 테스트
func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	//httptest 패키지의 NewRequest 함수로 테스트용 "/" 경로 요청 객체 req 만들기
	req := httptest.NewRequest("GET", "/", nil) // "/"경로 테스트

	mux := MakeWebHandler() //원 파일에서 만든 핸들러 인스턴스 가져와서 테스트
	mux.ServeHTTP(res, req) //핸들러 인스턴스의 ServeHTTP() 메소드 호출하여 요청 결과값 가져옴

	assert.Equal(http.StatusOK, res.Code)     //결과 코드 확인
	data, _ := io.ReadAll(res.Body)           // io.ReadAll()함수로 데이터([]byte) 읽어오기
	assert.Equal("Hello World", string(data)) //[]byte -> string, 일치 여부 확인
}

//"/bar" 경로 테스트
func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	//httptest 패키지의 NewRequest 함수로 테스트용 "/bar" 경로 요청 객체 req 만들기
	req := httptest.NewRequest("GET", "/bar", nil) //"/bar"경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
