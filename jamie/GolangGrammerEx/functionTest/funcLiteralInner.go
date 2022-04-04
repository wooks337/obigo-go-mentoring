package main

import "fmt"

func main() {
	i := 0 //함수 리터럴 외부에 있는 변수

	f := func() {
		i += 10 //함수 리터럴 내부에서 외부변수 i에 접근
	}
	i++

	f() //f 함수 타입 변수를 사용하여 함수 리터럴 실행

	fmt.Println(i)
}
