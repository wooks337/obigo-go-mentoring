package main

import (
	"fmt"
	"strings"
)

func main() {

	// ``백쿼드의 활용
	str1 := "hello\tworld"
	str2 := `hello\tworld`
	fmt.Println(str1)
	fmt.Println(str2)

	//문자열 길이 구하기
	str3 := "hello 한국!"
	runes := []rune(str3)
	fmt.Println("str3 len : ", len(str3))
	fmt.Println("runes len : ", len(runes))

	//String 합연산
	str4 := "abcdef"
	var str5 string = ""
	var strBuilder strings.Builder

	for _, v := range str4 {
		str5 += string(v)       //매번 메모리를 새로 만듬
		strBuilder.WriteRune(v) //기존 메모리 공간에 값을 채움
	}

}
