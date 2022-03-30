package hello

import "fmt"

type Hello struct {
}

func (h *Hello) PrintHello() string {
	fmt.Println("Hello world !!")
	return "hello"
}
