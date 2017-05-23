package three

import "testing"

func TestAdd11(t *testing.T) {
	r := add(1, 1)
	if r != 2 {
		t.Error("it should be 2 but was", r)
	}
}

func TestAdd12(t *testing.T) {
	r := add(1, 2)
	if r != 3 {
		t.Error("it should be 3 but was", r)
	}
}

func TestAdd22(t *testing.T) {
	r := add(2, 2)
	if r != 4 {
		t.Error("it should be 4 but was", r)
	}
}
