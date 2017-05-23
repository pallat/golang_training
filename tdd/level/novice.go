package main

import "fmt"

func step(a, b int) (int, int) {
	return b, (a + b)
}

func main() {
	a, b := 0, 1

	for i := 0; i < 15; i++ {
		fmt.Print(a, " ")
		a, b = step(a, b)
	}
}
