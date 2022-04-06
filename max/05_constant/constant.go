package main

import "fmt"

func main() {

	//기본선언
	const Name string = "Max"
	fmt.Println(Name)

	//iota를 이용한 선언
	const (
		Red   int = iota //0
		Blue             //1
		Green            //2
	)
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)

	//타입이 없는 상수
	const PI = 3.141592
	fmt.Println(PI * 100) //타입이 없기 때문에 오류가 안남

}
