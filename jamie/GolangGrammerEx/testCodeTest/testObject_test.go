package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestSquare1(t *testing.T) {
//	rst := square(9)
//	if rst != 81 {
//		t.Errorf("square(9) should be 81 but square(9) returns %d", rst)
//	}
//}
//
//func TestSquare2(t *testing.T) {
//	rst := square(3)
//	if rst != 9 {
//		t.Errorf("square(3) should be 9 but square(3) returns %d", rst)
//	}
//}

func TestSquare1(t *testing.T) {
	//assert 객체 생성 - 테스트 코드를 쉽게 만들 수 있는 다양한 메서드 포함
	assert := assert.New(t)
	//테스트 함수 호출 - Equal()메서드로 함수 출력 비교
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
