package main

import "fmt"

type yod func(str string) string

func yodFunc() yod {
	return func(s string) string {
		return fmt.Sprintf("My name is %s", s)
	}
}

func main() {
	fn := yodFunc()
	fmt.Println(fn("Roof"))

	var e func(string)
	e = echo
	e("testing")

	var p = func(s string) {
		fmt.Println(s)
	}

	func(s string) {
		fmt.Println(s)
	}("great")

	p("good")
}

func echo(s string) {
	fmt.Println(s)
}
