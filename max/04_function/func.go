package main

import "fmt"

func main() {

	i, i2, b := multiReturn(1, 2)
	fmt.Println(i, i2, b)

	fmt.Println(specifiedReturn(1, 2))
}

//멀티반환 함수
func multiReturn(a int, b int) (int, int, bool) {
	return a, b, true
}

//반환값 지정함수
func specifiedReturn(a int, b int) (c int, d int, e bool) {
	c, d = a, b
	e = true
	return
}
