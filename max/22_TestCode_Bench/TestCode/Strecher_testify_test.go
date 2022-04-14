package aaaa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	//assert.Equal(11, return10, "값이 같지않음")
	//assert.Greater(9, return10(), "값이 더 작음")
	//assert.NotNilf(returnNil(), "nil 값입니다")
	//assert.NotEqualf(10, return10(), "값이 같음")
	assert.Len("abcd", 5, "길이 다름")
}

func return10() int {
	return 10
}

func returnNil() interface{} {
	return nil
}
