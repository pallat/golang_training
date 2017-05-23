package three

import "testing"

func TestPartFi11(t *testing.T) {
	x, y := partFi(1, 1)
	if x != 1 {
		t.Error("it should be 1 but was", x)
	}
	if y != 2 {
		t.Error("it should be 2 but was", y)
	}

}

func TestPartFi12(t *testing.T) {
	x, y := partFi(1, 2)
	if x != 2 {
		t.Error("it should be 2 but was", x)
	}
	if y != 3 {
		t.Error("it should be 3 but was", y)
	}

}

func TestPartFi23(t *testing.T) {
	x, y := partFi(2, 3)
	if x != 3 {
		t.Error("it should be 3 but was", x)
	}
	if y != 5 {
		t.Error("it should be 5 but was", y)
	}

}

func TestPartFi35(t *testing.T) {
	x, y := partFi(3, 5)
	if x != 5 {
		t.Error("it should be 5 but was", x)
	}
	if y != 8 {
		t.Error("it should be 8 but was", y)
	}

}
