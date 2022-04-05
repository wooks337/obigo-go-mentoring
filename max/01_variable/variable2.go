package main

import "fmt"

func main() {

	//타입변환
	a := 3   //int
	b := 1.5 //float64

	//Error
	//var c int = b     //자동형변환 안됨 (같은 정수끼리여도 불가)
	//d := a * b        //다른타입끼리 연산 불가
	//var e int = b * 2 //자동형변환 안됨

	//형변환
	var c int = int(b)     //float64 -> int (소숫점 내림)
	d := float64(a) * b    //타입 통일
	var e int = int(b) * 2 // 1 * 2
	var f int = int(b * 2) // int(1.5 * 2)

	fmt.Print(c, d, e, f)

}
