package main

import "fmt"

func main() {
	m := make(map[int]int) //맵 생성
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)

	v1, ok1 := m[3]
	v2, ok2 := m[1]

	fmt.Println(v1, ok1)
	fmt.Println(v2, ok2)

	fmt.Println(m[3])
	fmt.Println(m[1])

}
