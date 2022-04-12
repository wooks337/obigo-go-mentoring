package main

import "fmt"

func main() {

	printResult(div(5, 3))
	printResult(div(5, 0))
}
func printResult(r float64, e error) {
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(r)
	}
}

func div(a, b int) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("0으로 나눌 수 없습니다")
	}
	return float64(a) / float64(b), nil
}
