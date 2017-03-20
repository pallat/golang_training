package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("....")
		fmt.Println("****")
		// if err := recover(); err != nil {
		// 	fmt.Println("recovered", err)
		// }
	}()

	panic("dead...")

	fmt.Println("go...")
}
