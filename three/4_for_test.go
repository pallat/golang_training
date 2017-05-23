package three

import "testing"

func TestFactorial(t *testing.T) {
	r := fact(5)
	if r != 120 {
		t.Error("5x4x3x2x1 = 120 but got", r)
	}
}
