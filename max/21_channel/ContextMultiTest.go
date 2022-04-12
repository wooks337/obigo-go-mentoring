package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)

	fmt.Println(cancel)
}
