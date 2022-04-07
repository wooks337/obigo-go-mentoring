package main

import "fmt"

func main() {

	s1 := structA{10}
	interA := StringerA(s1)
	interA.String()
	s2 := interA.(structA) //변환
	fmt.Println(s2.score)

	//다른형태로 변환하기
	//s5 := interA.(structB) //다른타입으로 변환해서 오류

	//주소로 받기
	s3 := &structA{11}
	interB := StringerA(s3)
	s4 := interB.(*structA)
	fmt.Println(s4.score)

}

type StringerA interface {
	String()
}

type structA struct {
	score int
}

func (s structA) String() {
}

type structB struct {
	score int
}

func (s structB) String() {
}
