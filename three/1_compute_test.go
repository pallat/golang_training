package three

import "testing"

func TestSubtract11(t *testing.T) {
	r := sub(1, 1)
	if r != 0 {
		t.Error("it should be 0 but was", r)
	}
}

func TestSubtract21(t *testing.T) {
	r := sub(2, 1)
	if r != 1 {
		t.Error("it should be 1 but was", r)
	}
}

func TestMultiply22(t *testing.T) {
	r := sub(2, 2)
	if r != 4 {
		t.Error("it should be 4 but was", r)
	}
}

func TestMultiply23(t *testing.T) {
	r := sub(2, 3)
	if r != 6 {
		t.Error("it should be 6 but was", r)
	}
}

func TestDivide63(t *testing.T) {
	r := sub(6, 3)
	if r != 2 {
		t.Error("it should be 2 but was", r)
	}
}

func TestDivide62(t *testing.T) {
	r := sub(6, 2)
	if r != 3 {
		t.Error("it should be 3 but was", r)
	}
}
