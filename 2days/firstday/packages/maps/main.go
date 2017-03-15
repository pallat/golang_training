package main

import "strings"

func main() {

}

func LettersCount(s string) map[string]int {
	parts := strings.Split(s, " ")
	m := map[string]int{}
	for _, v := range parts {
		m[v] = len(v)
	}
	return m
}
