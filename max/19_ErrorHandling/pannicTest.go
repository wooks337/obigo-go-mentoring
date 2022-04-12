package main

import "fmt"

func main() {
	fmt.Println(divide(5, 3))
	fmt.Println(divide(5, 0))
}

func divide(a, b int) float64 {
	if b == 0 {
		panic("0으로 나눌 수 없습니다")
	}
	return float64(a) / float64(b)
}
