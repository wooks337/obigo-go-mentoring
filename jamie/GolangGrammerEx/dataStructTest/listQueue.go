package main

import (
	"container/list"
	"fmt"
)

//Queue 구조체 정의
type Queue struct { //리스트를 이용한 Queue 구조체 정의
	v *list.List
}

//Push()메서드 정의
func (q *Queue) Push(val interface{}) { //(모든 타입의)요소 추가
	q.v.PushBack(val) //맨뒤에 요소 추가 - list의 PushBack()메소드 이용
}

//Pop()메서드 정의
func (q *Queue) Pop() interface{} { //Pop()메소드 - 맨 앞의 요소를 반환하고 삭제
	front := q.v.Front() //list의 Front()메소드 - 맨 앞의 요소 인스턴스 반환
	if front != nil {
		return q.v.Remove(front) //비어있지 않다면 Remove()메소드 호출하여 리스트 내 요소 삭제 후 반환
	}
	return nil //리스트가 비었으면 nil 반환
}

//NewQueue()함수 정의
func NewQueue() *Queue { // 새로운 Queue 인스턴스 - 내부 리스트 필드 list.New()함수로 초기화
	return &Queue{list.New()}
}
func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ { //NewQueue()함수에 for문으로 요소 추가 - 1~4
		queue.Push(i)
	}
	v := queue.Pop() //for문을 이용해서 리스트가 빌 때까지(v == nil) 요소를 하나씩 빼서 출력
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
