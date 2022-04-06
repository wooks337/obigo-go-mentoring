package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//부호가 있는 정수
	var int1 int8 = 5  //1byte	-2^7 ~ 2^7-1
	var int2 int64 = 5 //8byte	-2^63 ~ 2^63-1
	var int3 int = 5   //int64와 동일 , 8byte, 기본타입
	var int4 int
	int5 := 3 //선언 대입문을 통해 var와 타입 생략

	fmt.Println("int8 : ", unsafe.Sizeof(int1))
	fmt.Println("int64 : ", unsafe.Sizeof(int2))
	fmt.Println("int : ", unsafe.Sizeof(int3))
	fmt.Println("int의 기본값 : ", int4)
	fmt.Println("3의 기본 타입 : ", reflect.TypeOf(int5))
	fmt.Println()

	//부호가 없는 정수
	var uint1 uint8 = 5  //1byte 0 ~ 2^8-1
	var uint2 uint64 = 5 //8byte 0 ~ 2^64-1
	var uint3 uint = 5   //8byte 0 ~ 2^64-1
	fmt.Println("uint8 : ", unsafe.Sizeof(uint1))
	fmt.Println("uint64 : ", unsafe.Sizeof(uint2))
	fmt.Println("uint : ", unsafe.Sizeof(uint3))
	fmt.Println()

	//실수
	var float1 float32 = 0.1 //4byte
	var float2 float64 = 0.1 //8byte, 기본타입
	var float3 float64
	float4 := 0.5
	fmt.Println("float32 : ", unsafe.Sizeof(float1))
	fmt.Println("float64 : ", unsafe.Sizeof(float2))
	fmt.Println("float의 기본값 : ", float3)
	fmt.Println("0.5의 기본 타입 : ", reflect.TypeOf(float4))
	fmt.Println()

	//불리언
	var bool1 bool = true //1byte (cpu가 처리가능한 최소 크기는 1바이트)
	var bool2 bool
	bool3 := true
	fmt.Println("bool : ", unsafe.Sizeof(bool1))
	fmt.Println("bool의 기본값 : ", bool2)
	fmt.Println("true의 기본 타입 : ", reflect.TypeOf(bool3))
	fmt.Println()

	//문자열
	var str1 string = "abcdef" //영문1바이트
	var str2 string = "한글"     //한글3바이트
	var str3 string
	str4 := "abcde"

	fmt.Println("string (영문6글자) : ", len(str1))
	fmt.Println("string (한글2글자) : ", len(str2))
	fmt.Println("string 기본값 : ", str3) // ""임, nil 이 아님
	fmt.Println("\"abcde\"의 기본 타입 : ", reflect.TypeOf(str4))

}
