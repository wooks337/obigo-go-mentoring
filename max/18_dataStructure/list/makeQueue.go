package main

import (
	"container/list"
	"fmt"
)

func main() {

	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	var v any = 0
	for true {
		v = queue.Pop()
		if v == nil {
			break
		}
		fmt.Println(v)
	}
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	} else {
		return nil
	}
}
