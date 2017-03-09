package main

import "fmt"

func main() {
	h, w := helloAndWorld()
	fmt.Println(h, w)
}

func helloAndWorld() (string, string) {
	return "Hello", "world"
}
