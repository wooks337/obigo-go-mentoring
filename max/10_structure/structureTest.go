package main

import "fmt"

func main() {

	var student Student
	student.No = 15
	student.Name = "cat"
	student.Class = 3
	student.Score = 95

	fmt.Println(student)

	student2 := Student{"name", 2, 30, 100} //초기화
	fmt.Println(student2)

	student3 := Student{Name: "dog", Score: 76} //일부 초기화
	fmt.Println(student3)
}
