package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go square1(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) //채널을 닫아 데드락 해결
	wg.Wait()
}

func square1(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //채널에 데이터가 더 들어오기를 대기
		fmt.Println(n, "square = ", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
