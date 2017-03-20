package main

import "testing"

func TestPartialOfFibonacci(t *testing.T) {
	x, y := partFibonacci(1, 2)

	if x != 2 {
		t.Error("I need 2 but got", x)
	}

	if y != 3 {
		t.Error("I need 3 but got", y)
	}
}
