package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		stdin.ReadString('\n')
	}
	return num, err
}

func main() {
	for {
		fmt.Print("1~99 사이의 숫자를 입력하세요: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요!!")
		} else {
			fmt.Println("입력하신 숫자는 ", n, "입니다.")
		}
	}
}
