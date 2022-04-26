package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandlerGetList(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

func TestJsonHandlerGet(t *testing.T) {
	assert := assert.New(t)

	var student Student
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/5", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusNotFound, res.Code)
}

func TestHandlerPost(t *testing.T) {
	assert := assert.New(t) //import "github.com/stretchr/testify/assert"

	var student Student
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"id":0, "Name":"ccc", "Age":15, "Score":78}`)) //body데이터
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req) //Test서버 작동

	assert.Equal(http.StatusCreated, res.Code) //상태값 확인

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students",
		strings.NewReader(`{aaa}`))
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestHandlerDelete(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("bbb", list[0].Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("DELETE", "/students/5", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusNotFound, res.Code)
}
