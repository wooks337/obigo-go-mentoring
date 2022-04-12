package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	readEq("5 3")
	readEq("50 10")
}

func readEq(eq string) {
	result, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(result, "\n")
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { //err가 numError로 변환이 가능한가? (주소를 넣어줘야 함)
			fmt.Println("NumberError : ", numError)
		}
	}
}

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) //스캐너 생성
	scanner.Split(bufio.ScanWords)                      //한 단어씩 끊어 읽기

	pos := 0

	num1, len, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readSacn. pos : %d, err : %w", pos, err)
	}

	pos += len + 1
	num2, len, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readSacn. pos : %d, err : %w", pos, err)
	}

	return num1 * num2, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { //단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	num, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to conver word to int, word = %s, err = %w",
			word, err)
	}
	return num, len(word), nil
}
