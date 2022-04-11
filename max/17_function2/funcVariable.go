package main

import "fmt"

func main() {

	var operator func(int, int) int
	operator = getOperator("+")
	fmt.Println("결과 : ", operator(2, 3))

	var operator2 opFunc
	operator2 = getOperator("-")
	fmt.Println("결과 : ", operator2(2, 3))
}

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func getOperator(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "-":
		return minus
	default:
		return nil
	}
}

type opFunc func(int, int) int
