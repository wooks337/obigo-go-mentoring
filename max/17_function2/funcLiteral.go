package main

import "fmt"

func main() {

	fn := getOperator2("+")
	fmt.Println(fn(2, 3))

}

func getOperator2(op string) opFunc2 {
	switch op {
	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "-":
		return func(a int, b int) int {
			return a - b
		}
	default:
		return nil
	}
}

type opFunc2 func(int, int) int
