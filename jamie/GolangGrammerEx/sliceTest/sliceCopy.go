package main

import "fmt"

//slice1의 모든 요소값 slice2에 복사

//for문 이용
/*
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1)) // slice1과 같은 길이의 슬라이스 생성
	for i, v := range slice1 {
		slice2[i] = v
	}
*/

//append()를 이용
/*
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append([]int{}, slice1...)
*/
//copy()를 이용
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	slice1[1] = 100 // slice1의 요솟값을 변경해도 slice2는 다른 배열을 가리키므로 slice2의 요소는 변하지 않는다
	fmt.Println(slice1)
	fmt.Println(slice2)
}
