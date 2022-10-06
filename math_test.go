package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	result := Sum(1, 2)

	if result != 3 {
		t.Error("Invalid sum")
	}
}

func TestSub(t *testing.T)  {
	result := Sub(3, 2)

	if result != 1 {
		t.Error("Invalid sub")
	}
}

func TestDivide(t *testing.T) {
	result, _ := Divide(10, 2)

	if result != 5 {
		t.Error("Invalid division")
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 2)

	if result != 4 {
		t.Error("Invalid multiply")
	}
}
