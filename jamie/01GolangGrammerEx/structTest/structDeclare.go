package main

import "fmt"

type House struct { //House 구조체를 정의
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House              //House 구조체 변수를 선언
	house.Address = "경기도 성남시..." //각 필드 값 초기화
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소 :", house.Address)
	fmt.Printf("크기 : %d평\n", house.Size)
	fmt.Printf("가격 : %.2f억 원\n", house.Price)
	fmt.Println("타입 :", house.Type)
}
