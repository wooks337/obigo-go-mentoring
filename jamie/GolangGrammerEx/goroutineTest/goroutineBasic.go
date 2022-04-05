package main

import (
	"fmt"
	"time"
)

func PrintHangul() { //300ms 간격으로 가~사까지 출력
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}
func PrintNumbers() { //400ms 간격으로 1~5까지 출력
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

//3초간 대기 - 그동안 PrintHangul()이랑 PrintNumbers() 실행
//위 두 고루틴이 3초 이상 걸릴경우 메인함수는 도중에 종료됨
func main() {
	go PrintHangul()  //고루틴 생성
	go PrintNumbers() //고루틴 생성

	time.Sleep(3 * time.Second)
}
