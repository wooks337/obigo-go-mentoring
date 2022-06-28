package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

var f = func() { fmt.Println("예시 출력") }

func goCronsched_Clear() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Second().Do(func() {})
	s.Every(1).Minute().Do(func() {})

	fmt.Println(len(s.Jobs())) // Job 개수 출력
	s.Clear()                  // 모든 Job Clear

	fmt.Println(len(s.Jobs())) // Clear 후 Job 개수 출력
	s.StartAsync()
}

func goCronsched_Cron() {
	s := gocron.NewScheduler(time.UTC)

	s.Cron("*/1 * * * *").Do(func() { fmt.Println("every min") }) // 매 분
	//s.Cron("0 1 * * *").Do(func() {})                             // 매일 오전 1시
	//s.Cron("0 0 * * 6,0").Do(func() {})                           // 주말 업데이트

	s.StartAsync()
}

func goCronsched_Rand() {
	s := gocron.NewScheduler(time.UTC)
	s.EveryRandom(1, 5).Seconds().Do(func() {
		fmt.Println("랜덤 알람!")
	})

	s.StartAsync()
}

func goCronsched_JobUpdate() {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every("1s").Do(func() { //초기 job 생성하고 시작
		fmt.Println("job 예제")
	})
	s.StartAsync()

	time.Sleep(time.Second * 5)

	s.Job(j).Every("1m").Update() //기존 job 호출해서 update 시간단위
	time.Sleep(time.Minute * 10)
}

func goCronsched_JobLen() {
	s := gocron.NewScheduler(time.UTC)
	s.Every("1s").Do(func() {})
	s.Every("2s").Do(func() {})
	s.Every("3s").Do(func() {})

	fmt.Println(len(s.Jobs())) //--스케줄러의 job 목록의 길이
	fmt.Println(s.Len())       //--스케줄러의 Job 개수 리턴
}

func goCronsched_LessJob() {
	s := gocron.NewScheduler(time.UTC)

	s.Every("1s").Do(func() {})
	s.Every("2s").Do(func() {})
	s.Every("3s").Do(func() {})
	s.StartAsync()
	fmt.Println(s.Jobs())
	fmt.Println(s.Less(0, 1))
	// 인덱스가 어떻게 정해지는지는 모르겠지만
	// 각 job의 인덱스가 있고 second 변수의 job이 first 변수 job보다 늦게 실행되면 true를 반환한다
}

func goCronsched_NextRun() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("10:30").Do(func() {})
	s.Every(1).Day().At("08:30").Do(func() {})
	s.StartAsync()
	_, t := s.NextRun()
	fmt.Println(t.Format("13:24")) //--다음 job 실행 시각 포맷
}

func goCronsched_Remove() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Week().Do(f)
	s.StartAsync()
	s.Remove(f)
	fmt.Println(s.Len())

}

func goCronsched_Tag() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Second().Tag("tag").Do(func() {
		fmt.Println("5초 tag")
	})
	s.Every(10).Seconds().Tag("tag").Do(func() {
		fmt.Println("10초 tag")
	})

	s.RunByTag("tag") //---지정 태그 서치하여 전부 실행
	//s.RunByTagWithDelay("tag", time.Second*10) //---지정 태그 서치 후 10초 뒤 실행
	s.StartAsync()
	time.Sleep(time.Minute) //1분동안 실행~
}

func goCronsched_MaxConcurrentJobs() {
	s := gocron.NewScheduler(time.UTC)

	//한번에 몇개의 job을 실행할 지 정하는 메소드
	s.SetMaxConcurrentJobs(1, gocron.RescheduleMode)
	s.Every(1).Seconds().Do(func() {
		fmt.Println("매 초 실행으로 스케줄링 되어 있지만 Max concurrent가 설정되어 있어서 5초 마다 실행")
		time.Sleep(5 * time.Second)
	})
	s.StartAsync()
	time.Sleep(time.Minute)
}

func goCronsched_Start() {
	s := gocron.NewScheduler(time.UTC)
	specificTime := time.Date(2022, time.June, 28, 17, 10, 0, 0, time.UTC)

	s.Every(1).Day().At("17:10").Do(func() { //왜 인지 얘는 시간을 맞춰도 저시간에 출력이 안된다...
		fmt.Println("이 시간에 알람!")
	})
	s.Every(1).Minute().StartAt(specificTime).Do(func() {
		fmt.Println("현재 utc 시각 출력")
	})
	s.Every(1).Minute().StartImmediately().Do(func() {
		fmt.Println("바로 시작!", time.Now())
	})

	s.StartAsync()
	time.Sleep(time.Minute * 5)
}

func goCronsched_Stop() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Second().Do(func() {})

	s.StartAsync()
	fmt.Println(s.IsRunning())

	s.Stop()
	fmt.Println(s.IsRunning())
}

func main() {
	//goCronsched_Clear()
	//goCronsched_Cron()
	//goCronsched_Rand()
	//goCronsched_JobUpdate()
	//goCronsched_JobLen()
	//goCronsched_LessJob()
	//goCronsched_NextRun()
	//goCronsched_Remove()
	//goCronsched_Tag()
	//goCronsched_MaxConcurrentJobs()
	//goCronsched_Start()
	goCronsched_Stop()
}
