package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	var echo = func(m string) {
		println(m)
	}
	echo("test")
}
