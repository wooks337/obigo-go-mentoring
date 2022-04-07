package main

import (
	"fmt"
	"sort"
)

func main() {

	//일반 정렬
	slice := []int{4, 3, 5, 1, 2}
	sort.Ints(slice)
	fmt.Println(slice)

	//구조체 정렬

	s := []Student{
		{"나", 3},
		{"다", 5},
		{"가", 1}}

	sort.Sort(Students(s))
	fmt.Println(s)
}

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
