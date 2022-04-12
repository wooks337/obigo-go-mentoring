package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	wg2.Add(1)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go PrintEverySecond2(ctx)

	wg2.Wait()
}

func PrintEverySecond2(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-ctx.Done(): //context의 cancle로 인해 실행
			wg2.Done()
			return
		}
	}
}
