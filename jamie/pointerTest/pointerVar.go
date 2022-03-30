package main

import "fmt"

func main() {
	var a int = 500
	var p *int //int 포인터 변수 p 선언

	p = &a //a 변수의 메모리 주소를 p 변수의 값으로 대입

	fmt.Printf("p의 값: %p\n", p)            //p의 메모리 주소값 출력
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) //p가 가리키는 메모리의 값

	*p = 100
	fmt.Printf("a의 값: %d\n", a) //p로 인해 변경된 a의 값
}
