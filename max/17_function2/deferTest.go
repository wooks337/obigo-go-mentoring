package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Create Error : ", err)
	}
	defer fmtTestFunc()
	defer f.Close()
	defer fmt.Println("222 파일을 닫았습니다.")

	fmt.Println("333 파일에 Hello World 작성")
	fmt.Fprintln(f, "Hello World")
}

func fmtTestFunc() {
	fmt.Println("111 반드시 호출되는 fmt")
}
