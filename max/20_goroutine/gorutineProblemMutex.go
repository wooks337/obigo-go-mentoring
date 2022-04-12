package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mutex sync.Mutex
	a := 0

	wg.Add(2)
	go minusMutex(&a, &wg, &mutex)
	go plusMutex(&a, &wg, &mutex)
	wg.Wait()

	fmt.Println("a = ", a)
}

func minusMutex(a *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	fmt.Println("minusMutex 시작")
	mutex.Lock()
	fmt.Println("minusMutex Lock시작")
	defer mutex.Unlock()
	for i := 0; i < 10000; i++ {
		*a--
	}
	fmt.Println("minusMutex 종료")
	wg.Done()
}

func plusMutex(a *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	fmt.Println("plusMutex 시작")
	mutex.Lock()
	fmt.Println("plusMutex Lock시작")
	defer mutex.Unlock()

	for i := 0; i < 10000; i++ {
		*a++
	}
	fmt.Println("plusMutex 종료")
	wg.Done()
}
