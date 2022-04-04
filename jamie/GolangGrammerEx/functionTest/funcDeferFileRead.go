package main

import (
	"fmt"
	"os"
)

func Helloworld() {
	file, err := os.Open("text.txt") //파일 열기
	defer file.Close()               //defer 파일 닫기

	if err != nil { //파일열떄 에러 발생 검증
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)

	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}

func main() {
	Helloworld()
	fmt.Println("Done")
}
