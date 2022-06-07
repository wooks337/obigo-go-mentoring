package main

import "fmt"

const M = 10 //나머지 연산의 분모

func hash(d int) int {
	return d % M //나머지 연산
}
func main() {
	m := [M]int{} //값 저장용 배열 생성 - 배열길이가 10, 요소타입은 int, 요소는 없음

	m[hash(23)] = 10  //키:23 값:10
	m[hash(259)] = 50 //키:259 값:50

	fmt.Printf("%d = %d\n", 23, m[hash(23)])   //m[3] = 23
	fmt.Printf("%d = %d\n", 259, m[hash(259)]) //m[9] = 259

	m[hash(33)] = 50 //m[hash(33)] = m[3]이기 때문에 m[hash(23)]과 해시 충돌 발생
	//인덱스 위치마다 값대신 리스트를 저장하여 문제 해소

}
