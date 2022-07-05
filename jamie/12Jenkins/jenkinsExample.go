package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func runCron() {
	s := gocron.NewScheduler(time.UTC)
	s.Every("5s").Do(func() {
		fmt.Println("5초마다 반복")
	})

	s.StartAsync()
	time.Sleep(time.Second * 20)
}

func main() {
	runCron()
}
