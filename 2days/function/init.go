package main

import e "hi/echo"
import "fmt"

func init() {
	fmt.Println("initial...")
}

func main() {
	e.Say()
}
