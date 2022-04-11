package main

import "fmt"

func main() {
	CaptureLoop1()
	CaptureLoop2()
}

func CaptureLoop1() {
	f := make([]func(), 3) //함수 리터럴 3개의 슬라이스
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Printf("값 = %d, 주소 = %p\n", i, &i) //외부변수 i를 사용하기 때문에 i의 값이 아닌 i의 주소를 참조해서 값을 가져옴
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3) //함수 리터럴 3개의 슬라이스
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i //i의 값을 v로 복사, => v는 반복문 실행할때마다 새로 생성
		f[i] = func() {
			fmt.Printf("값 = %d, 주소 = %p\n", v, &v) //v의 주소는 반복문 실행할때마다 달라짐
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}
