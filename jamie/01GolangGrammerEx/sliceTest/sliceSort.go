package main

import (
	"fmt"
	"sort"
)

//int 슬라이스 정렬
/*
func main() {
	slice := []int{5, 2, 7, 4, 1}
	sort.Ints(slice)

	fmt.Println(slice)
}
*/
//구조체 슬라이스 (나이순) 정렬

type Student struct {
	Name string
	Age  int
}

//[]Student의 별칭 타입 Students
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }

func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}
