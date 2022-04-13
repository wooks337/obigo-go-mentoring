package main

import "fmt"

func main() {

	slice := []int{5, 4, 3, 2, 1}

	fmt.Println(slice)
	bubbleSort(slice)
	fmt.Println(slice)
}

func bubbleSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			changeWhenBig(&slice[j], &slice[j+1])
		}
	}
}

func changeWhenBig(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}
