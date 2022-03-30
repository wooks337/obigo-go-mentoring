package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	//range를 이용하여 모든 배열 요소 순회
	//인덱스(i)와 요소값(v) 반환
	for i, v := range t {
		fmt.Printf("[%d: %v]\n", i, v)
	}
	for _, v := range t {
		fmt.Print(v, "\t")
	}
}
