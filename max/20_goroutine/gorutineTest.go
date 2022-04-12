package main

import (
	"fmt"
	"time"
)

func main() {
	go PrintA()
	go PrintB()
	time.Sleep(2 * time.Second)
}

func PrintA() {
	for i := 0; i < 3; i++ {
		fmt.Print("A ")
		time.Sleep(100 * time.Millisecond)
	}
}
func PrintB() {
	for i := 0; i < 3; i++ {
		fmt.Print("B ")
		time.Sleep(100 * time.Millisecond)
	}
}
