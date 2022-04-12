package main

import (
	"fmt"
)

func main() {
	f1()
	fmt.Println("프로그램이 계속 실행")
}

func f1() {
	fmt.Println("함수 시작")
	defer func() {
		if r := recover(); r != nil {
			//if r, ok := recover().(net.Error); ok {
			//	fmt.Println("panic 복구 : ", r)
			//}
			fmt.Println("panic 복구 : ", r)
		}
	}()

	f2()
	fmt.Println("함수 종료")
}

func f2() {
	fmt.Println("9 / 3 = ", f3(9, 3))
	fmt.Println("9 / 0 = ", f3(9, 0))
}

func f3(a, b int) float64 {
	if b == 0 {
		panic("0으로 나눌 수 없습니다")
	}
	return float64(a) / float64(b)
}
