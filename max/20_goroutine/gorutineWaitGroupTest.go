package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3) //작업 갯수 설정

	go SumAtoB(1, 1000000000)
	go SumAtoB(1, 1000000000)
	go SumAtoB(1, 1000000000)
	wg.Wait() //종료 대기
}

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(a, "부터 ", b, "까지의 합은 ", sum)
	wg.Done() //종료 선언
}
