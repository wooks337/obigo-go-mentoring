package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200 // array2의 2번 인덱스 요소 200으로 변경
}
func changeSlice(slice2 []int) {
	slice2[2] = 200
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("배열: ", array)
	fmt.Println("슬라이스: ", slice)
}

// 배열은 새로 대입될 때 배열 전체가 복사되며 원본과 다른 메모리 주소를 가지게 된다.
// 슬라이스는 새로 대입 시 메모리 주소와 len, cap 모두 복사되며 원본과 같은 메모리 주솟값을 가진다.
