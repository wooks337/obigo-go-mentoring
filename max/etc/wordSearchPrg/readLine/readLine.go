package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("dataMax.txt")
}

func PrintFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
