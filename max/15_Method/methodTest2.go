package main

import "fmt"

func main() {

	var a myInt = 10
	fmt.Println(a.add(20))
	a.add2() //원래는 변환하여 사용해야 하지만 go가 알아서 해줌
}

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func (a *myInt) add2() {

}
