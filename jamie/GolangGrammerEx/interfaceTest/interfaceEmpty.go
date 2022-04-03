package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println(int(t))
	case float64:
		fmt.Println(float64(t))
	case string:
		fmt.Println(string(t))
	default:
		//그 외 타입인 경우 타입과 값 출력
		fmt.Printf("타입: %T, 값: %v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(101)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}
