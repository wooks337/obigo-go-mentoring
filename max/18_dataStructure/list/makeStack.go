package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := NewStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	var v any = 0
	for v != nil {
		v = stack.Pop()
		if v == nil {
			break
		}
		fmt.Println(v)
	}
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}
func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	} else {
		return nil
	}

}
