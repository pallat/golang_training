package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case fmt.Stringer:
		print(v.String())
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Println(i.(bool), v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type a struct {
	s string
	v string
}

func (a a) String() string {
	return a.s
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(a{"test", "try"})
}
