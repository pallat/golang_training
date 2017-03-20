package main

import "fmt"

func main() {
	var world = func() string {
		return "world"
	}
	fmt.Println("Hello", world())
}
