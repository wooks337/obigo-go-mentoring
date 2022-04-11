package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["A"] = "Apple"
	m["B"] = "Banana"
	m["C"] = "Canada"
	m["C"] = "Change"

	delete(m, "B")

	fmt.Println("m[\"A\"] : ", m["A"])
	fmt.Println("m[\"A\"] : ", m["B"])
	fmt.Println("m[\"C\"] : ", m["C"])

	v, ok := m["A"]
	fmt.Println(v, ", ", ok)
}
