package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//MultipleFromString() 함수 선언 - 매개변수: string 타입 str - 반환타입: int, 에러
//readNextInt() 함수는 변환된 숫자, 읽은 글자 수, 에러를 반환값으로 받음
//readNextInt() 함수를 호출하여 err 발생 시, readNextInt() 함수의 에러를 감싸서 pos 정보를 에러에 추가함
func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) //스캐너 생성
	scanner.Split(bufio.ScanWords)                      //한 단어씩 끊어서 읽기

	pos := 0                          //pos번째 글자
	a, n, err := readNextInt(scanner) //a = 첫번째 단어가 변환된 숫자
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err) // 에러 감싸기
	}
	pos += n + 1                      //pos번째 글자
	b, n, err := readNextInt(scanner) //b = 두번째 단어가 변환된 숫자
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err) // 에러 감싸기
	}
	return a * b, nil
}

//readNextInt() 함수 선언 - 매개변수: Scanner(구조체) 타입 scanner의 메모리 주소를 가리키는 포인터 변수 - 반환타입: int, int, 에러
//다음 단어를 읽어서 숫자로 변환하여 변환된 숫자, 읽은 글자 수, 에러를 반환합니다.

//if Scan()메소드 !scanner이면  -> 스캔 실패 메세지 출력
//아니면 scanner로 읽어온 단어 = word
//Atoi() 함수로 문자열을 int로 변형시켜 int, error 반환
//	if 에러 발생 시(Atoi() 함수 결과에서 숫자가 아닌 문자가 섞여 있을 경우), Atoi 변형 실패 에러 메세지 출력, NumberError 타입 에러 반환
//	아니면 변환된 숫자, 읽은 글자 수, nil(에러)을 반환
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err)
		//이때 에러를 %w로 감싸서 에러 뿐만 아니라 word와 같은 다른 정보도 같이 하나의 에러로 반환할 수 있게 에러를 감싼다.
	}
	return number, len(word), nil
}

//
func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
