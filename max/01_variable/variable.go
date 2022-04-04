package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//부호가 있는 정수
	var int1 int8 = 5  //1byte	-2^7 ~ 2^7-1
	var int2 int64 = 5 //8byte	-2^63 ~ 2^63-1
	var int3 int = 5   //int64와 동일 , 8byte

	fmt.Println("int8 : ", unsafe.Sizeof(int1))
	fmt.Println("int64 : ", unsafe.Sizeof(int2))
	fmt.Println("int : ", unsafe.Sizeof(int3))

	//부호가 없는 정수
	var uint1 uint8 = 5  //1byte 0 ~ 2^8-1
	var uint2 uint64 = 5 //8byte 0 ~ 2^64-1
	var uint3 uint = 5   //8byte 0 ~ 2^64-1
	fmt.Println("uint8 : ", unsafe.Sizeof(uint1))
	fmt.Println("uint64 : ", unsafe.Sizeof(uint2))
	fmt.Println("uint : ", unsafe.Sizeof(uint3))

	//문자열
	var float1 float32 = 0.1 //4byte
	var float2 float64 = 0.1 //8byte
	fmt.Println("float32 : ", unsafe.Sizeof(float1))
	fmt.Println("float64 : ", unsafe.Sizeof(float2))

	//불리언
	var bool1 bool = true //1byte (cpu가 처리가능한 최소 크기는 1바이트)
	fmt.Println("bool : ", unsafe.Sizeof(bool1))

	//문자열
	var str1 string = "abcdef" //영문1바이트
	fmt.Println("string (영문6글자) : ", len(str1))

	var str2 string = "한글" //한글3바이트
	fmt.Println("string (한글2글자) : ", len(str2))

}
