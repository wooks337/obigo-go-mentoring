package main

import "fmt"

// 별칭 타입
type myInt int

//myInt 별칭 타입을 리시버로 갖는 add() 메서드
func (a myInt) add(b int) int {
	return int(a) + b
}
func main() {
	var a myInt = 10
	fmt.Println(a.add(30))
	var b int = 20                //myInt와 int는 다른 타입!!!
	fmt.Println(myInt(b).add(50)) //따라서 int b는 myInt타입으로 타입변환을 하고 add()메서드 사용 가능
}
