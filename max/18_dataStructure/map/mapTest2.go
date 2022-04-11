package main

import "fmt"

func main() {

	m := make(map[int]Product)

	m[16] = Product{"볼펜", 1500}
	m[45] = Product{"지우개", 300}
	m[356] = Product{"자", 700}

	for k, v := range m {
		fmt.Println("key = ", k, ", value = ", v)
	}

}

type Product struct {
	Name  string
	Price int
}
