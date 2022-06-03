package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //취소 가능한 컨텍스트 생성
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second) //5초 후에
	cancel()                    // 취소 함수 호출 -> Done()채널에 시그널 보내기

	wg.Wait() //모든 작업이 끝날 때까지 대기
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): //Done()채널에서 취소 명령 확인
			wg.Done() //시그널 수신하고 프로그램 종료
			return
		case <-tick:
			fmt.Println("Tick")

		}
	}
}
