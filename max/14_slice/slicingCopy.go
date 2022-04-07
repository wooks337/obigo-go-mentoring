package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	//슬라이스 값만 복사 V1
	slice2 := make([]int, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}

	//슬라이스 값만 복사 V2
	slice3 := make([]int, 3, 10) //len3, cap10
	cnt := copy(slice3, slice1)  //slice3에 len 만큼 slice1을 덮는다
	fmt.Println(slice3)
	fmt.Println(cnt)

	//슬라이스 값만 복사 V3
	slice4 := make([]int, len(slice1))
	cnt2 := copy(slice4, slice1)
	fmt.Println(slice4)
	fmt.Println(cnt2)
}
