package main

import "fmt"

func main() {
	fmt.Println("합 : ", sum(1, 2, 3, 4, 5))
}

func sum(nums ...int) int { //가변인수함수

	sum := 0

	fmt.Printf("nums의 타입 : %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}
