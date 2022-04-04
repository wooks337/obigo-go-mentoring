package main

import "fmt"

type Data struct { //Data 타입 구조체
	value int
	data  [200]int
}

func ChangeData(arg Data) { //arg 데이터 변경
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data1 Data

	ChangeData(data1)
	fmt.Printf("value= %d\n", data1.value)
	fmt.Printf("data[100] = %d\n", data1.data[100])
}
