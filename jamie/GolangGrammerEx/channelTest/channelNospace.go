package main

import "fmt"

func main() {
	var ch chan int = make(chan int)

	ch <- 9
	fmt.Println("Never Print")
}
