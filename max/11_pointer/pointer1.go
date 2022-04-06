package main

import "fmt"

func main() {

	//기본
	var a = 5
	var p *int
	p = &a

	*p = 10
	fmt.Println("a의 주소 : ", &a)
	fmt.Println("a의 값 : ", a)
	fmt.Println("p의 주소 : ", p)
	fmt.Println("p의 값 : ", *p)

	var p2 = new(int) //메모리 생성
	fmt.Println("p2의 주소 : ", p2)
}
