package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println(slice)

	slice = append(slice, 5, 6, 7, 13, 5)
	fmt.Println(slice)

	for i := 6; i < 16; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
