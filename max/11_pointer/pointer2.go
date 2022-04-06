package main

import "fmt"

func main() {

	var student Student

	changeStudent1(student) //값복사
	fmt.Println(student)

	changeStudent2(&student) //주소복사
	fmt.Println(student)

	var pStudent *Student = &Student{}
	changeStudent2(pStudent) //주소복사
	fmt.Println(*pStudent)

	var pStudent2 = new(Student)
	changeStudent2(pStudent2) //주소복사
	fmt.Println(*pStudent2)

}

func changeStudent2(std *Student) {
	std.Age = 10
	std.Name = "hello10"
}

func changeStudent1(std Student) {
	std.Age = 10
	std.Name = "hello10"
}

type Student struct {
	Age  int
	Name string
}
