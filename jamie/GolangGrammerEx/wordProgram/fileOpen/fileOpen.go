package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

func PrintFile(filename string) {
	file, err := os.Open(filename) //파일 열기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return
	}
	defer file.Close() //함수 종료 전 파일 닫기

	scanner := bufio.NewScanner(file) // 스캐너 생성, 한 줄씩 읽기
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //스캐너로 읽은 내용 출력하기
	}
}