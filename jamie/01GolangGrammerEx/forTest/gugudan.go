package main

import "fmt"

func main() {
	dan := 2
	x := 1

	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, x, dan*x)
			x++
			if x == 10 {
				break
			}
		}
		x = 1
		dan++
		if dan == 10 {
			break
		}
	}
	fmt.Println("구구단이 종료되었습니다.")
}
