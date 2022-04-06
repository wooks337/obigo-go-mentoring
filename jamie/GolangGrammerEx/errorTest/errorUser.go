package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) { //float 타입의 f 매개변수를 가지고 float와 error를 반환값으로 가지는 Sqrt함수
	if f < 0 {
		return 0, fmt.Errorf(
			"제곱근은 양수여야 합니다. f:%g", f) //f가 음수면 에러 반환
	}
	return math.Sqrt(f), nil
}
func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err) //에러 출력
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}

//출력
//Error: 제곱근은 양수여야 합니다. f:-2
