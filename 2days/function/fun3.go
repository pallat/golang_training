package main

import "fmt"

func main() {
	fn := makeFunc()
	fn("Hello, World")
}

func makeFunc() func(m string) {
	return func(m string) {
		fmt.Println(m)
	}
}
