package main

import (
	"container/list"
	"fmt"
)

//Stack 구조체 정의
type Stack struct { //리스트를 이용한 Stack 구조체 정의
	v *list.List
}

//Push()메소드 정의
func (s *Stack) Push(val interface{}) { //(모든 타입의)요소 추가
	s.v.PushBack(val) //맨뒤에 요소 추가
}
func (s *Stack) Pop() interface{} { //Pop()메소드 - 맨 앞의 요소를 반환하고 삭제
	back := s.v.Back() //Back()메소드 이용하여 맨 뒤 요소부터 반환
	if back != nil {
		return s.v.Remove(back) //비어있지 않다면 Remove()메소드 호출하여 리스트 내 요소 삭제 후 반환
	}
	return nil //리스트가 비었으면 nil 반환
}

func NewStack() *Stack { // 새로운 Stack 인스턴스 - 내부 리스트 필드 list.New()함수로 초기화
	return &Stack{list.New()}
}

func main() {
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
