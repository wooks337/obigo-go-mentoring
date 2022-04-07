package main

import "fmt"

func main() {

	//슬라이스 중간 요소값 추가 V1
	slice := []int{1, 2, 4, 5, 6}
	idx := 2 //2번인덱스에 값 추가 예정

	slice = append(slice, 0)
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	slice[idx] = 3
	fmt.Println(slice)

	//슬라이스 중간 요소값 추가 V2
	slice1 := []int{1, 2, 4, 5, 6}
	idx = 2
	slice1 = append(slice1[:idx], append([]int{3}, slice1[idx:]...)...)
	fmt.Println(slice1)

	//슬라이스 중간 요소값 추가 V3
	slice2 := []int{1, 2, 4, 5, 6}
	idx = 2
	slice2 = append(slice2, 0)
	copy(slice2[idx+1:], slice2[idx:])
	slice2[idx] = 3
	fmt.Println(slice2)

}
