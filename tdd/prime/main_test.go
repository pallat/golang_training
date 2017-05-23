package main

import "testing"

func TestCanDiv(t *testing.T) {
	b := canDiv(6, 3)
	if !b {
		t.Error(b)
	}
	b = canDiv(6, 4)
	if b {
		t.Error(b)
	}
}
