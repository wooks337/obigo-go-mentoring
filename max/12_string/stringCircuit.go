package main

import "fmt"

func main() {

	str4 := "hello 월드!"
	//문자열 순회 1.인덱스 사용 -> 1바이트씩 가져오기 때문에 한글 못읽음
	for i := 0; i < len(str4); i++ {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", str4[i], str4[i], str4[i])
	}
	fmt.Println("========================")

	//문자열 순회 2.rune[] 변환 -> 4바이트씩 가져오므로 잘 됨, but 불편, 메모리 낭비
	runes2 := []rune(str4)
	for i := 0; i < len(runes2); i++ {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", runes2[i], runes2[i], runes2[i])
	}
	fmt.Println("========================")
	//문자열 순회 3.range 사용
	for _, v := range str4 {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", v, v, v)
	}
	fmt.Println("========================")

}
