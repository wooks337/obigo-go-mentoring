package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}
	for _, arr := range a { // arr은 a배열
		for _, v := range arr { // v는 a배열의 각 요소
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
