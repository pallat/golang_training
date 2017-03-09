package main

import "fmt"

func main() {
	fn := func() func(m string) {
		return func(m string) {
			fmt.Println(m)
		}
	}()
	fn("Hello, World")
}
