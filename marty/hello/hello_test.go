package hello

import (
	"testing"
)

func TestHello_PrintHello(t *testing.T) {
	hello := Hello{}
	result := hello.PrintHello()

	if result == "" {
		t.Error("Empty data....")
	}
}
