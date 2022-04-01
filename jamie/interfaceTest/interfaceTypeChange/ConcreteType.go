package main

import "fmt"

type Stringer interface {
	String() string //string 타입의 String()메서드
}
type Student struct { //Student 구조체
	Age int
}

func (s *Student) String() string { //Student 타입의 String() 메서드
	return fmt.Sprintln(s.Age)
}

func PrintAge(stringer Stringer) {

	s := stringer.(*Student) //*Student 타입으로 변환
	fmt.Println(s.Age)
}
func main() {
	s := &Student{15} //*Student 타입의 변수 s 선언 및 초기화
	PrintAge(s)
}
