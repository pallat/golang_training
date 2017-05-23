package three

import "testing"

func TestSwap(t *testing.T) {
	x, y := swap("hello", "world")
	if x != "world" {
		t.Error("it should be world but was", x)
	}
	if y != "hello" {
		t.Error("it should be hello but was", y)
	}
}
