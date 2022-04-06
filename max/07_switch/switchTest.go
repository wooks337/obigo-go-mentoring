package main

import "fmt"

func main() {
	switch age := getMyAge(); true {
	case age > 150, age < 0:
		fmt.Println("너무 많거나 적음")
	case age >= 20 && age < 30:
		fmt.Println("20대")
	default:
		fmt.Println("그 외")
	}
}

func getMyAge() int {
	return 28
}
