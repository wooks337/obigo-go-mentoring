package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("text.txt") //파일 생성
	if err != nil {                 //에러 확인
		fmt.Println("Failed to create a file")
		return
	}
	defer fmt.Println("defer문은 반드시 호출됩니다.") //지연 수행 코드
	defer f.Close()                         //지연 수행 코드
	defer fmt.Println("파일을 닫았습니다.")         //지연 수행 코드

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World") //파일에 텍스트 작성
}
