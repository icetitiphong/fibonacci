package Fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {

	given := 10
	want := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	get := Fibonacci(given)

	if len(want) != len(get) {
		t.Errorf("want %q but got %q\n", want, get)
	}
}
