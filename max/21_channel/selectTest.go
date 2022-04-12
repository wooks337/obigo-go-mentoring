package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square2(&wg, ch, quit)

	for i := 0; i < 5; i++ {
		ch <- i * 2
	}

	quit <- true

	wg.Wait()
}

func square2(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for true {
		select { //ch와  quit 모두를 기다림
		case n := <-ch:
			fmt.Println("Square : ", n*n)
			time.Sleep(time.Second)
		case <-quit:
			fmt.Println("종료")
			wg.Done()
			return
		}
	}
}
