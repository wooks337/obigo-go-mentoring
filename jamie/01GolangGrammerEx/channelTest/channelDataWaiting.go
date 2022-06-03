package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	//ch에 데이터가 들어오기를 기다렸다가 데이터가 들어오면 n에 값 복사하여 for문 실행
	for n := range ch { //채널이 닫히면 자동종료
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() //실행 x -> for range문은 계속 데이터가 들어오기를 무한 대기하여 wg.Done()이 실행되지 않음
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 //채널에 데이터 10번 입력
	}
	close(ch) //채널 닫음!!!
	wg.Wait() //작업 완료 대기
}
