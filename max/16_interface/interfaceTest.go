package main

import "fmt"

func main() {

	student := Student{"철수", 12}
	var stringer Stringer

	stringer = student
	fmt.Println(stringer.ToString())
}

type Stringer interface {
	ToString() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) ToString() string {
	return fmt.Sprintln(s.Name, ",", s.Age)
}
