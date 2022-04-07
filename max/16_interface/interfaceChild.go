package main

import "fmt"

func main() {

	s1 := struct1{} //모든 인터페이스에서 사용가능
	s2 := struct2{} //childA 에만 가능
	s3 := struct3{} //childB 에만 가능
	s4 := struct4{} //모든 인터페이스에서 사용불가

	p := parent(s1)

	ca1 := childA(s1)
	ca2 := childA(s2)

	cb1 := childB(s1)
	cb2 := childB(s3)

	fmt.Print(s4, p, ca1, ca2, cb1, cb2)
}

type childA interface {
	aaa()
	xxx()
}

type childB interface {
	bbb()
	xxx()
}

type parent interface {
	childA
	childB
}

type struct1 struct {
}

func (s struct1) aaa() {
}
func (s struct1) bbb() {
}
func (s struct1) xxx() {
}

type struct2 struct {
}

func (s struct2) aaa() {
}
func (s struct2) xxx() {
}

type struct3 struct {
}

func (s struct3) bbb() {
}
func (s struct3) xxx() {
}

type struct4 struct {
}

func (s struct4) aaa() {
}
func (s struct4) bbb() {
}
