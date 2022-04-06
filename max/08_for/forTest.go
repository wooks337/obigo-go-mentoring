package main

import "fmt"

func main() {

	//기본선언
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//중첩 for문
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//중첩 for문 한번에 나가기
OuterFor:
	for i := 0; i < 5; i++ {
		for j := 0; j <= 5; j++ {
			for k := 0; k <= 5; k++ {
				if i == 1 && j == 1 && k == 1 {
					fmt.Println("종료")
					break OuterFor
				} else {
					println("i = ", i, ", j = ", j, ", k = ", k)
				}
			}
		}
		fmt.Println()
	}

}
