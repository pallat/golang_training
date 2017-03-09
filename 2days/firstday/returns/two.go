package main

import "fmt"

func main() {
	fmt.Println(helloAndWorld())
}

func helloAndWorld() (string, string) {
	return "Hello", "world"
}
