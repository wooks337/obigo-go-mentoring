package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil) //"/students" 경로에 대한 GET 요청 만들기
	//students 경로 테스트
	mux := MakeWebHandler() //웹 핸들러 요청
	mux.ServeHTTP(res, req) //결과 반환 확인

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list) //결과 변환
	assert.Nil(err)                                //결과 확인
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil) //id가 1인 학생 테스트

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)

	assert.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil) //id가 2인 학생 테스트

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("bbb", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":0, "Name":"ccc", "Age":15, "Score":78}`)) //새로운 학생 데이터으 json 포맷을 post 요청에 포함

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code) //응답 코드검사

	//새로 추가된 학생데이터 GET 메소드로 확인
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)
	//추가된 학생 데이터
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)
	//id = 1 삭제 요청

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code) //응답 코드 검사

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("bbb", list[0].Name)
}
