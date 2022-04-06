package main

import "fmt"

func main() {

	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr1 : ", arr1)

	//인덱스 접근
	for i := 0; i < len(arr1); i++ {
		arr1[i] = arr1[i] * 10
	}
	fmt.Println(arr1)

	//for range
	for i, v := range arr1 { //필요없을 경우 _ 로 대체
		fmt.Println("index : ", i, ", value : ", v)
	}

	//배열 복사
	arr2 := arr1 //값 복사
	fmt.Println("arr2 = ", arr2)
	arr2[0] = 100
	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)

	//이차원배열
	arr3 := [2][5]int{{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}}
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[0]); j++ {
			fmt.Print(arr3[i][j], " ")
		}
		fmt.Println()
	}

	for _, vArr := range arr3 {
		for _, v := range vArr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
