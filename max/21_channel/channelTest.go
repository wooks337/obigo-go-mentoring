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
	go square(&wg, ch)

	time.Sleep(time.Second) //1초 대기 후 채널에 값 넣어줌
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch //값이 들어올때까지 대기

	fmt.Println(n, " square = ", n*n)
	wg.Done()
}
