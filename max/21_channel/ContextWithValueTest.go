package main

import (
	"context"
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func main() {
	wg3.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	go square4(ctx)
	wg3.Wait()
}

func square4(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Println("Square : ", n*n)
	}
	wg3.Done()
}
