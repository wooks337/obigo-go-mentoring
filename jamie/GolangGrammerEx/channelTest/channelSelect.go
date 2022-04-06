package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { //ch와 quit 양쪽을 모두 기다린다
		//ch 채널이 먼저 시도되고 ch 채널에서 데이터를 읽어 10개의 제곱이 출력된다
		//quit 채널에 서 데이터를 읽어 square함수를 종료한다.
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool) //quit과 ch 채널을 만든다

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true
	wg.Wait()
}
