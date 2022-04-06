package main

import "fmt"

func main() {

	//기본선언
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	arr3 := [5]int{1: 1, 3: 3}
	fmt.Println(arr3)
}
