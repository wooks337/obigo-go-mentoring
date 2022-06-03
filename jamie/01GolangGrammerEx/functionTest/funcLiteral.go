package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		//함수 리터럴을 이용하여 더하기 함수 정의 및 반환
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		//함수 리터럴을 이용하여 곱하기 함수 정의 및 반환
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4) //함수타입 변수를 사용하여 함수 호출
	fmt.Println(result)
}
