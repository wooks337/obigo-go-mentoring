package BenchmarkCode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkFibonacciRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursion(20)
	}
}

func BenchmarkFibonacciBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciBasic(20)
	}
}

func TestFibonacciRecursion(t *testing.T) {

	assert := assert.New(t)

	assert.Equal(0, fibonacciRecursion(-1), "fiboRecursion(-1) 은 0이 나와야 함")
	assert.Equal(0, fibonacciRecursion(0), "fiboRecursion(0) 은 0이 나와야 함")
	assert.Equal(1, fibonacciRecursion(1), "fiboRecursion(1) 은 0이 나와야 함")
	assert.Equal(2, fibonacciRecursion(3), "fiboRecursion(2) 은 0이 나와야 함")
	assert.Equal(233, fibonacciRecursion(13), "fiboRecursion(13) 은 233이 나와야 함")
}

func TestFibonacciBasic(t *testing.T) {

	assert := assert.New(t)

	assert.Equal(0, fibonacciBasic(-1), "fibonacciBasic(-1) 은 0이 나와야 함")
	assert.Equal(0, fibonacciBasic(0), "fibonacciBasic(0) 은 0이 나와야 함")
	assert.Equal(1, fibonacciBasic(1), "fibonacciBasic(1) 은 0이 나와야 함")
	assert.Equal(2, fibonacciBasic(3), "fibonacciBasic(2) 은 0이 나와야 함")
	assert.Equal(233, fibonacciBasic(13), "fibonacciBasic(13) 은 233이 나와야 함")
}

func fibonacciRecursion(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

func fibonacciBasic(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	one := 1
	two := 0
	result := 0

	for i := 2; i <= n; i++ {
		result = one + two
		two = one
		one = result
	}
	return result
}
