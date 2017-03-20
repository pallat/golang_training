package main

import "fmt"

func main() {
	fmt.Println(swap("Hello","World"))
}

func swap(x,y string) (string,string) {
	return y,x
}
