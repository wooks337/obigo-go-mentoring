package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	a := 0

	wg.Add(2)
	go minus(&a, &wg)
	go plus(&a, &wg)
	wg.Wait()

	fmt.Println("a = ", a)
}

func minus(a *int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		*a--
	}
	fmt.Println("minus 종료")
	wg.Done()
}

func plus(a *int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		*a++
	}
	fmt.Println("plus 종료")
	wg.Done()
}
