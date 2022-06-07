package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스 f
	fmt.Println("ValueLoop")

	//for 루프 내에서 i 변수를 캡쳐한 함수 리터럴을 f 슬라이스 값으로 저장
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i) // 캡쳐된 i값 출력
		}
	}
	//저장된 f 슬라이스 각 항목을 호출
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3) // 함수 리터럴 3개를 가진 슬라이스 f
	fmt.Println("ValueLoop2")

	//i값을 저장하는 변수를 새로 만들어 새로 만든 변수를 캡쳐
	for i := 0; i < 3; i++ {
		v := i // v 변수에 i 값 복사
		f[i] = func() {
			fmt.Println(v) // 캡쳐된 v값 출력
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop() //3,3,3
	// i 변수 캡쳐 시, 캡쳐하는 순간의 i값이 복사되는 것이 아님
	// i 변수 자체가 참조
	// 따라서 for문 종료 후 최종 i의 값인 3이 함수 리터럴 호출 시 i값이 됨
	CaptureLoop2() //0,1,2
	// v변수 선언 후 i값 복사 -> 함수리터럴에서 v변수 캡쳐
	// for문 내부에서 선언된 v값은 for문 반복시 마다 새로운 v 변수 생성
	// f 슬라이스의 각 함수 리터럴 요소는 서로 다른 v 변수를 캡쳐
}
