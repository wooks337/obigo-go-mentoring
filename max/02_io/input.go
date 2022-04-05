package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Print("값 입력 : ")
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("입력갯수 : ", n, ", a = ", a, ", b = ", b)
	}

	//fmt.Print("값 입력 : ")
	//n, err := fmt.Scanf("%d %d", &a, &b)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("입력갯수 : ", n, ", a = ", a, ", b = ", b)
	//}

	//fmt.Print("값 입력 : ")
	//n, err := fmt.Scanln(&a, &b)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("입력갯수 : ", n, ", a = ", a, ", b = ", b)
	//}

}
