package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {
	//t.Fail()
	//fmt.Println("Hello from our test")
	var (
		expected = 101
		a        = 1
		b        = 0
	)

	have := calculateSum(a, b)

	if have != expected {
		t.Errorf("expected %d but have %d", expected, have)
	}
}
