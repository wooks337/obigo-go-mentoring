//재귀함수
package main

import "fmt"

func printNo(n int) {
	if n == 0 { //재귀호출 탈출 조건
		return
	}
	fmt.Println(n)

	printNo(n - 1)          //재귀호출
	fmt.Println("After", n) //재귀호출 이후 출력
}

func main() {
	printNo(3) //함수호출
}
