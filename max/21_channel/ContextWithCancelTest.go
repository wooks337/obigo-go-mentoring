package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	wg1.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) //인수로 상위 컨텍스트를 넣어줌
	//상위 컨텍스트가 없다면 기본 컨텍스트인 context.Background()를 넣음
	//cancel을 반환받아 언제든 종료 가능
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() //5초후에 cancel을 통해 Done()채널에 시그널을 보냄

	wg1.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-ctx.Done(): //context의 cancle로 인해 실행
			wg1.Done()
			return
		}
	}
}
