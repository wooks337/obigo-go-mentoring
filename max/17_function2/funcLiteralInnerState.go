package main

import "fmt"

func main() {
	i := 0

	f := func() {
		fmt.Println("i 값 = ", i)
		i += 10
	}

	i++
	f()
	fmt.Println(i)
}
