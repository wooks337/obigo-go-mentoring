package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)

	go square3(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

func square3(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second) //1초 이후 시그널,
	//time.Tick() 일정 시간 주기로 신호를 보내주는 채널 생성
	terminate := time.After(10 * time.Second) //10초 이후 시그널
	//time.After() 일정 시간 후에 신호를 보내주는 채널 생성
	for true {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square : ", n*n)
			time.Sleep(time.Second)
		}
	}
}
