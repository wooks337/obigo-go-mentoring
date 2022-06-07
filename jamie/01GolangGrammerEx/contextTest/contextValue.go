package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	//컨텍스트에 값 "9" 추가
	go square(ctx)

	wg.Wait()
}
func square(ctx context.Context) {
	//Value()메서드로 키 값을 읽어오기
	//Value()메서드 반환값 : 빈 인터페이스
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
