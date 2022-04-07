package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci1(t *testing.T) {
	assert := assert2.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(3) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}
func TestFibonacci2(t *testing.T) {
	assert := assert2.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(3) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) { //재귀호출
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) { //반복문
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}

//반복문을 이용한 방식이 재귀호출 방식보다 빠르다!!
