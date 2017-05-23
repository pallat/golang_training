package captcha

import "testing"

func TestFirstPatternOnePlusOne(t *testing.T) {
	c := &captcha{
		pattern:       1,
		firstOperand:  1,
		seconfOperand: 1,
		operator:      1,
	}

	s := c.String()
	if s != "1 + one" {
		t.Error("'1 + one' is expected but got", s)
	}
}

func TestSecondPatternOnePlusOne(t *testing.T) {
	c := &captcha{
		pattern:       2,
		firstOperand:  1,
		seconfOperand: 1,
		operator:      1,
	}

	s := c.String()
	if s != "one + 1" {
		t.Error("'one + 1' is expected but got", s)
	}
}

func TestFirstPatternOneMinusOne(t *testing.T) {
	c := &captcha{
		pattern:       1,
		firstOperand:  1,
		seconfOperand: 1,
		operator:      2,
	}

	s := c.String()
	if s != "1 - one" {
		t.Error("'1 - one' is expected but got", s)
	}
}

func TestFirstPatternOneMultiplyOne(t *testing.T) {
	c := &captcha{
		pattern:       1,
		firstOperand:  1,
		seconfOperand: 1,
		operator:      3,
	}

	s := c.String()
	if s != "1 * one" {
		t.Error("'1 * one' is expected but got", s)
	}
}
