package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	checked = "true"
	x       int
	y       int
)

func main() {
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Println(partFibonacci(partFibonacci(partFibonacci(x, y))))
}
