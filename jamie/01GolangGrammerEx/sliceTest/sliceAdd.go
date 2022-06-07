package main

import "fmt"

//for문 이용
/*
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	//맨 뒤에 요소 추가
	slice = append(slice, 0)
	idx := 2 //추가하려는 위치

	//맨 뒤부터 추가 위치까지 값을 옮긴다
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	slice[idx] = 100
	fmt.Println(slice)
}
*/
//append() 이용
/*
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	fmt.Println(slice)
}
*/
//copy() 이용
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	//[1,2,3,4,5,6,0] <- [2,3,4,5,6] = [1,2,2,3,4,5,6]

	slice[idx] = 100
	fmt.Println(slice)
}
