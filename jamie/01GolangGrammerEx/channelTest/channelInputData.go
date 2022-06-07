package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) //채널 생성

	wg.Add(1)
	go square(&wg, ch) //square함수를 실행하는 고루틴 1개 생성 -> 이제 square() 함수는 main 루틴이 아닌 새로운 고루틴에서 main과 동시에 실행된다
	ch <- 9            //채널에 데이터 주입
	wg.Wait()          //작업 완료 대기
}

func square(wg *sync.WaitGroup, ch chan int) { //waitGroup 객체의 메모리 주소와 채널 인스턴스를 매개변수로 받는 square 함수
	n := <-ch // 채널 인스턴스에서 데이터 빼오려고 시도, 들어올때 까지 대기

	time.Sleep(time.Second)         //1초 대기
	fmt.Printf("Square: %d\n", n*n) //square : 81
	wg.Done()
}
