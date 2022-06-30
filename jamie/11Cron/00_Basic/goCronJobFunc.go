package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func goCronjob_Error() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("bad time")
	j := s.Jobs()[0]
	fmt.Println(j.Error())
}

func goCronjob_IsRunning() {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every(10).Seconds().Do(func() {
		time.Sleep(time.Second * 2)
	})
	fmt.Println(j.IsRunning()) //--false

	s.StartAsync()

	time.Sleep(time.Second)
	fmt.Println(j.IsRunning()) //--true

	time.Sleep(time.Second)
	s.Stop()

	time.Sleep(time.Second)
	fmt.Println(j.IsRunning()) //--false
}

func goCronjob_LastRun() {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Second().Do(func() {})
	s.StartAsync()

	fmt.Println("Last run:", job.LastRun())
}

func goCronjob_LimitRunsTo() {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every(1).Second().LimitRunsTo(5).Do(func() {
		fmt.Println("run limit")
	})
	//s.StartBlocking() //--이걸로 하면 run limit 5번 출력되는데 그 이후로 대기 상태 들어가서 fmt.print 출력 안함
	s.StartAsync()
	//ch := make(chan bool)
	//ch <- true
	//close(ch)
	//<-ch

	fmt.Println("LimitRunsTo()RunCount() : ", j.RunCount())
}

func goCronjob_ScheduledAtTime() {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every(1).Day().At("13:50").Do(func() {})
	s.StartAsync()
	fmt.Println("ScheduledAtTime() : ", j.ScheduledAtTime()) //--해당 job이 실행될 구체적 시각 출력
	fmt.Println("ScheduledTime() : ", j.ScheduledTime())     //--Job의 다음 run 시각
}

func main() {
	goCronjob_Error()
	goCronjob_IsRunning()
	goCronjob_LimitRunsTo()
	goCronjob_LastRun()
	goCronjob_ScheduledAtTime()

}
