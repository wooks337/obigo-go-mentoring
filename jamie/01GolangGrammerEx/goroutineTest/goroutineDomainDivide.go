package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface { //Do()메서드를 가지는 job 인터페이스
	Do()
}
type SquareJob struct { //SquareJob 구조체
	index int
}

func (j *SquareJob) Do() { //SquareJob 구조체의 메모리 주소를 타입으로 갖는 Do() 메서드
	fmt.Printf("%d 작업시작\n", j.index) //1초 대기후 제곱값 만들기 작업
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업완료 - 결과: %d\n", j.index, j.index*j.index)
}
func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ { //10가지 작업 할당
		jobList[i] = &SquareJob{i}
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		job := jobList[i] //각 작업을 고루틴으로 실행
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
