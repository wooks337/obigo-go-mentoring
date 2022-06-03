package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice1[1] = 100
	fmt.Println(slice1, slice2)

	slice1 = append(slice1, 500)
	fmt.Println(slice1, slice2)
}
