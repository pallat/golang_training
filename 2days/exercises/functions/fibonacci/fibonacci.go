package main

import "fmt"

func main() {
	a, b := 0, 1

	for i := 0; i < 30; i++ {
		fmt.Println(a)
		a, b = partFibonacci(a, b)
	}
}
