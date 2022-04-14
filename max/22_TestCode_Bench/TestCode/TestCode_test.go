package aaaa

import (
	"testing"
)

func TestSquare2(t *testing.T) {
	result := square(3)
	if result != 5 {
		t.Errorf("오류발생")
	}
}

func TestSquare(t *testing.T) {
	result := square(9)
	if result != 81 {
		t.Errorf("오류발생")
	}
}

func square(x int) int {
	return x * x
}
