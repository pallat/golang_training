package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(subtract(42, 13))
	fmt.Println(multiply(42, 13))
	fmt.Println(divide(42, 12))
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y float32) float32 {
	return x / y
}
