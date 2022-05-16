package Controller

import (
	"encoding/json"
	"net/http"
	"obigo-go-mentoring/jamie/GormRestAPI/database"
	"sort"
)

var students map[int]database.Student

type Students []database.Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].StudentID < s[j].StudentID
}

//전체 학생 데이터 조회
func GetStudents(w http.ResponseWriter, req *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)                       //상태 값
	w.Header().Set("Content-Type", "application/json") //header 값 설정
	json.NewEncoder(w).Encode(list)                    //import "encoding/json"
}
