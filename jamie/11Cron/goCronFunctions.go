package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func goCron_Error() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("bad time")
	j := s.Jobs()[0]
	fmt.Println(j.Error())
}

func goCron_IsRunning() {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every(10).Seconds().Do(func() {
		time.Sleep(time.Second * 2)
	})
	fmt.Println(j.IsRunning())

	s.StartAsync()

	time.Sleep(time.Second)
	fmt.Println(j.IsRunning())

	time.Sleep(time.Second)
	s.Stop()

	time.Sleep(time.Second)
	fmt.Println(j.IsRunning())
}

func goCron_LastRun() {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Second().Do(func() {})
	s.StartAsync()

	fmt.Println("Last run:", job.LastRun())
}

func goCron_LimitRunsTo() {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every(1).Second().LimitRunsTo(5).Do(func() {
		fmt.Println("run limit")
	})
	//s.StartAsync() //--이걸로 run하면 runCount 1로 나오고
	s.StartBlocking() //--이걸로 하면 run limit 5번 출력되는데 그 이후로 대기 상태 들어가서 fmt.print 출력 안함
	fmt.Println(j.RunCount())
}

func main() {
	goCron_Error()
	goCron_IsRunning()
	goCron_LastRun()
	goCron_LimitRunsTo()
}
