package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)             //WaitGroup객체 생성; 고루틴 2개
	fork := &sync.Mutex{} //포크와 수저 뮤텍스
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저") //A는 포크 먼저
	go diningProblem("B", spoon, fork, "수저", "포크") //B는 수저 먼저
	wg.Wait()                                      //모든 작업 종료까지 대기
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s가 밥을 먹으려고 합니다.\n", name)
		first.Lock() //첫번째 뮤텍스 획득 시도
		fmt.Printf("%s가 %s 획득\n", name, firstName)
		second.Lock() //두번째 뮤텍스 획득 시도
		fmt.Printf("%s가 %s 획득\n", name, secondName)

		fmt.Printf("%s가 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
