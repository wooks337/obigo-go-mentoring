package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3] //{2,3}, len=2, cap=4

	fmt.Println(arr)
	fmt.Println(slice)

	arr[1] = 20
	fmt.Println(arr)
	fmt.Println(slice)

	slice = append(slice, 400)
	fmt.Println(arr)
	fmt.Println(slice)

	//cap 조절
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3:4] //{2,3}, len=2, cap=3 (1번인덱스에서 4번인덱스 까지 3개)
	fmt.Println(slice2)
}
