package main

import "fmt"

func f() { //panic 복구
	fmt.Println("f()함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g() //호출 순서 : main() -> f() -> g() -> h()
	fmt.Println("f() 함수 끝")
}

func g() { //h() 함수 호출
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int { //panic 발생
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() { //f() 함수 호출
	f()
	fmt.Println("프로그램이 게속 실행됨")
}
