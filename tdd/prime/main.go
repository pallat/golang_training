package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 2; i < 101; i++ {
		fmt.Println(i, prime(i))
	}
}

func prime(n int) string {
	var can bool
	for i := n - 1; i > 0; i-- {
		can = can || canDiv(n, i)
	}

	if can {
		return strconv.Itoa(n)
	}
	return "Prime"
}

func canDiv(n, i int) bool {
	return (n % i) == 0
}
