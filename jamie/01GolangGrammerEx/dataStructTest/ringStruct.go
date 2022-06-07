package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) //요소가 5개인 링 생성, r = 현재위치 가리키는 포인터
	n := r.Len()     //링의 길이 반환

	for i := 0; i < n; i++ { //순회하면서 모든 요소에 값 대입 - A B C D E
		r.Value = 'A' + i
		r = r.Next()
	}
	for j := 0; j < n; j++ { //순회하면서 값 출력 - A B C D E
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println() // 한줄 띄우고

	for j := 0; j < n; j++ { //역순하면서 값 출력 - 현재위치 A 이후로 E D C B
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}

}
