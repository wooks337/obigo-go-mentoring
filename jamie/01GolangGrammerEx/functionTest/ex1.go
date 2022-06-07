package main

import "fmt"

func Multiple(a int, b int) int {
	return a * b
}

func main() {
	c := Multiple(3, 4)
	fmt.Println(c)
}
