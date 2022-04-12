package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg1.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg1.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Println(name, " 밥을 먹으려 합니다")
		first.Lock()
		fmt.Println(name, " ", firstName, "획득")
		second.Lock()
		fmt.Println(name, " ", secondName, "획득")

		fmt.Println(name, "밥을 먹습니다")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg1.Done()
}
