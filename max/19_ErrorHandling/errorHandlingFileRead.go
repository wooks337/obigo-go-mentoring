package main

import (
	"bufio"
	"fmt"
	"os"
)

const FileName string = "dataMax.txt"

func main() {
	file, err := ReadFile(FileName)
	if err != nil {
		err = WriteFile(FileName, "hello world")
		if err != nil {
			fmt.Println("파일 생성 실패, ", err)
			return
		}
		file, err = ReadFile(FileName)
		if err != nil {
			fmt.Println("파일 읽기 실패", err)
			return
		}
	}
	fmt.Println("파일 내용 : ", file)
}

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	readString, _ := reader.ReadString('\n')
	return readString, nil
}

func WriteFile(fileName string, line string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}
