package main

import "fmt"

func main() {
	m := make(map[string]string) //맵 생성

	m["이화랑"] = "서울시 광진구" //키와 값 추가
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구" //값 변경

	fmt.Printf("송하나의 주소 : %s\n", m["송하나"]) //값 출력
	fmt.Printf("최번개의 주소 : %s\n", m["최번개"])

	for k, v := range m { //맵 순회하여 값 출
		fmt.Println(k, v)
	}

}
