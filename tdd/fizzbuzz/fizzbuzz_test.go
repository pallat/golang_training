package fizzbuzz

import "testing"

func TestInputOneShouldBeOne(t *testing.T) {
	r := fizzbuzz(1)
	if r != "1" {
		t.Error("it should be 1 but was", r)
	}
}

func TestInputTwoShouldBeTwo(t *testing.T) {
	r := fizzbuzz(2)
	if r != "2" {
		t.Error("it should be 2 but was", r)
	}
}

func TestInputThreeShouldBeFizz(t *testing.T) {
	r := fizzbuzz(3)
	if r != "Fizz" {
		t.Error("it should be Fizz but was", r)
	}
}

func TestInputFourShouldBeFour(t *testing.T) {
	r := fizzbuzz(4)
	if r != "4" {
		t.Error("it should be 4 but was", r)
	}
}

func TestInputFiveShouldBeBuzz(t *testing.T) {
	r := fizzbuzz(5)
	if r != "Buzz" {
		t.Error("it should be Buzz but was", r)
	}
}

func TestInputSixShouldBeFizz(t *testing.T) {
	r := fizzbuzz(6)
	if r != "Fizz" {
		t.Error("it should be Fizz but was", r)
	}
}

func TestInputTenShouldBeBuzz(t *testing.T) {
	r := fizzbuzz(10)
	if r != "Buzz" {
		t.Error("it should be Buzz but was", r)
	}
}

func TestInputFifteenShouldBeFizzBuzz(t *testing.T) {
	r := fizzbuzz(15)
	if r != "FizzBuzz" {
		t.Error("it should be FizzBuzz but was", r)
	}
}
