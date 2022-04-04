package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { // 함수
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // 메소드 선언
	a.balance -= amount
}

func main() {
	a := &account{100} //balance가 100인 account 포인터 변수 a
	withdrawFunc(a, 30)
	a.withdrawMethod(30)

	fmt.Println(a.balance)
}
