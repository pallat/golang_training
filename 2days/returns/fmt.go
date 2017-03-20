package hello

import "fmt"

// Println display string
func Println(s string) {
	i, err := fmt.Println("Hello,", s)
	fmt.Println(i, err)
}
