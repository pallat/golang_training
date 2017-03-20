package main

import "fmt"

func main() {
	fmt.Println(helloAndWorld())
}

func helloAndWorld() (h string, w string) {
	h = "Hello"
	w = "world"
	return
}
