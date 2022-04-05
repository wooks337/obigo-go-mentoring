package main

import "fmt"

func main() {

	a := 10
	b := 324.12345
	c := "문자열"

	//출력포맷
	fmt.Print("a : ", a, ", b : ", b, ", c : ", c)
	fmt.Println("a : ", a, ", b : ", b, ", c : ", c)
	fmt.Printf("a : %d, b : %f, c : %s \n", a, b, c)
	fmt.Printf("a : %v, b : %v, c : %v \n", a, b, c)
	fmt.Printf("a : %T, b : %T, c : %T \n", a, b, c)

	//너비
	fmt.Println("======정수자릿수지정=======")
	fmt.Printf("%%d    : %d 종료\n", a)
	fmt.Printf("%%5d   : %5d 종료\n", a)
	fmt.Printf("%%05d  : %05d 종료\n", a)
	fmt.Printf("%%-05d : %-05d 종료\n", a)

	fmt.Println("======실수자릿수지정=======")
	fmt.Printf("%%f     : %f 종료\n", b)     //소숫점6자리
	fmt.Printf("%%g     : %g 종료\n", b)     //소숫점 끝까지, 지수
	fmt.Printf("%%08.2f : %08.2f 종료\n", b) //8자리, 소숫점2자리, 왼족0
	fmt.Printf("%%08.2g : %08.2g 종료\n", b) //8자리, 소숫점2자리, 왼쪽0, 지수
	fmt.Printf("%%8.2g  : %8.2g 종료\n", b)  //8자리, 소숫점2자리, 왼쪽 공백, 지수
}
