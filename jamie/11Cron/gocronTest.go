package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

//1시간 동안
func runBasicCron() {
	s := gocron.NewScheduler(time.UTC)
	s.Every("1m").Do(func() {
		log.Println("Every single minute!!")
	})

	s.StartAsync()               //-- 내부 go routine을 가지고 있음
	time.Sleep(time.Minute * 60) //go routine exit을 방지
}

func runWorldClock() {
	nyc, _ := time.LoadLocation("America/Los_Angeles")
	//seoul, _ := time.LoadLocation("South Korea/Seoul")

	s := gocron.NewScheduler(time.UTC)
	s.ChangeLocation(nyc)
	fmt.Println(s.Location())

	s.Cron("6 5 * * *").Do(func() {
		log.Println("뉴욕은 정각입니다")
	})
	s.StartAsync()
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	//runBasicCron()
	runWorldClock()
}
