package main

import (
	"errors"
)

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Divide(a int, b int) (int, error) {
	if b < 0 {
		return 0, errors.New("b should be bigger than 0")
	}
	return a / b, nil
}

func Multiply(a int, b int) (int) {
	return a * b
}
