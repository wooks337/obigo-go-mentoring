package main

import (
	"container/ring"
	"fmt"
)

func main() {

	r := ring.New(5) //ring의 현재위치

	for i := 0; i < r.Len(); i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%c \n", r.Value)
		r = r.Next()
	}

}
