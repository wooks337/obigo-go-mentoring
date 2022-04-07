package main

import "fmt"

func main() {

	var slice1 = []int{1, 2, 3}       //1,2,3
	var slice2 = []int{1, 3: 2, 5: 4} //1,0,0,2,0,4
	var slice3 = make([]int, 3)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	//요소 추가
	var slice4 []int
	for i := 0; i < 5; i++ {
		slice4 = append(slice4, i)
	}
	fmt.Println(slice4)

	var slice5 []int
	slice5 = append(slice5, 0, 1, 2, 3, 4)
	fmt.Println(slice5)

	//같은 배열 참조
	slice6 := make([]int, 3, 5)
	fmt.Println(slice6)
	slice7 := append(slice6, 1, 2) //남은공간이 2개였는데 2개를 추가했기 때문에
	//slice6과 slice7은 같은 배열을 가리키게 된다
	slice6[1] = 1
	fmt.Println(slice6)
	fmt.Println(slice7)
}
