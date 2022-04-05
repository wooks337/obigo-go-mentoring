package main

import "math"

const epsilon = 0.000001

func equal1(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func equal2(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {

	//실수의 비교연산
	a := 0.1
	b := 0.2
	c := 0.3

	println("a + b == c : ", a+b == c)
	println(a + b)

	println("a + b equal1 c : ", equal1(a+b, c))

	d := 0.0000000000002
	e := 0.0000000000001
	f := 0.0000000000003

	println("d + e equal2 f : ", equal2(d+e, f))

}
