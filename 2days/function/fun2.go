package main

import "fmt"

func main() {
	fn := echo
	fn("Hello, World")
}

func echo(m string) {
	fmt.Println(m)
}
