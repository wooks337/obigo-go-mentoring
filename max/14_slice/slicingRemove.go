package main

import "fmt"

func main() {

	//슬라이스 중간 요소값 삭제 V1
	slice := []int{1, 2, 3, 4, 5}
	idx := 2 //2번인덱스 삭제 예정
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice1 := slice[:len(slice)-1] //마지막 인덱스 제거
	fmt.Println(slice1)

	//슬라이스 중간 요소값 삭제 V2
	slice2 := []int{1, 2, 3, 4, 5}
	idx = 2 //2번인덱스 삭제 예정
	slice3 := append(slice2[:idx], slice2[idx+1:]...)
	fmt.Println(slice3)

}
