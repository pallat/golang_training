package fizzbuzz

import (
	"strconv"
)

func fizzbuzz(number int) string {
	if (number % 15) == 0 {
		return "FizzBuzz"
	}
	if (number % 3) == 0 {
		return "Fizz"
	}
	if (number % 5) == 0 {
		return "Buzz"
	}
	return strconv.Itoa(number)
}
