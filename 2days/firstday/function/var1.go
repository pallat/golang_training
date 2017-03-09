package main

import "fmt"

var echo = func(m string) {
	fmt.Print(m)
}

func main() {
	echo("test string")
}
