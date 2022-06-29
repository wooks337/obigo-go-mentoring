package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func runCron() {

	//새로운 스케줄러 생성
	s := gocron.NewScheduler(time.UTC)

	//cron Job생성
	s.Every(30).Seconds().Do(func() {
		fmt.Println("30초 마다 반복")
	})
	//--문자열로 시간 표시
	s.Every("1m").Do(func() {
		fmt.Println("1분 마다 반복")
	})
	//--시간 설정
	s.Every(1).Day().At("10:55").Do(func() {
		fmt.Println("매일 아침 9시 50분에 출력")
	})
	s.Every(1).Wednesday().At("10:55:00").Do(func() {
		fmt.Println("수요일 아침 9시 50분에 출력")
	})
	//--cron 표현식 지원
	s.Cron("*/2 * * * *").Do(func() {
		fmt.Println("cron 2분 마다 반복")
	})
	//--Tag로 job 실행
	s.TagsUnique()
	s.Every(1).Minute().Tag("foo").Do(func() {
		fmt.Println("foo 태그로 실행")
	})
	s.RunByTag("foo")

	//스케줄러 실행 방법
	//1 -- 비동기로 실행
	s.StartAsync()
	//2 -- 스케줄러를 실행하고 현재 실행 경로 차단
	//s.StartBlocking()

	for {
		time.Sleep(time.Second)
	}
}

//1시간 동안
func runBasicCron() {
	s := gocron.NewScheduler(time.UTC)
	s.Every("1m").Do(func() {
		log.Println("Every single minute!!")
	})

	s.StartAsync()               //-- 내부 go routine을 가지고 있음
	time.Sleep(time.Minute * 60) //go routine exit을 방지
}

//세계도시 정각마다 알림
func runWorldlock() {
	nyc, _ := time.LoadLocation("America/New_York")

	s := gocron.NewScheduler(nyc)
	//fmt.Println(s.Location())
	//s.Cron("*/1 * * * *").Do(func() {
	//	log.Println("뉴욕, 1분마다 입니다")
	//})

	s.Cron("0 * * * *").Do(func() {
		log.Println("뉴욕, 정각 입니다")
	})
	s.Cron("CRON_TZ=Europe/Paris 0 * * * *").Do(func() {
		log.Println("파리, 정각 입니다")
	})

	s.StartAsync()
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	runCron()
	runBasicCron()
	runWorldlock()
}
