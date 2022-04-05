package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) //입력을 읽는 객체

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') //줄바꿈 문자가 나올때까지 읽어서 비움
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
