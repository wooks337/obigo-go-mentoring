package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //WaitGroup 객체 생성 -> 두 함수에서 전역변수로 사용

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합계는 %d입니다.\n", a, b, sum)
	wg.Done() //작업 완료 표시
}

func main() {
	wg.Add(10) //총 작업 개수 설정 -> 고루틴 10개 만들기
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 10000)
	}
	wg.Wait() //모든 작업이 완료되길 기다림.
	//wg.Done()메서드가 작업개수만큼 호출될 때까지 기다렸다가, 잔여 작업개수가 0이 되면, wg.Wait() 메서드 종료 후, 다음 줄로 넘어감
	fmt.Println("모든 계산이 완료되었습니다.")
}
