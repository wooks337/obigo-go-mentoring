package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
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
	_, _ = s.Every(1).Day().At("10:25").Do(func() {
		fmt.Println("매일 아침 9시 30분에 출력")
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
	//s.StartAsync()
	//2 -- 스케줄러를 실행하고 현재 실행 경로 차단??
	s.StartBlocking()
}

func main() {
	runCron()
}
