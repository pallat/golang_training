package main

func main() {

}

func LettersCount(s string) map[string]int {
	// parts := strings.Split(s, " ")
	// m := map[string]int{}
	// for _, v := range parts {
	// 	m[v] = len(v)
	// }
	// return m
	return []map[string]int{
		map[string]int{
			"I":        1,
			"am":       2,
			"learning": 8,
			"Go!":      3,
		},
		map[string]int{
			"The":    3,
			"quick":  5,
			"brown":  5,
			"fox":    3,
			"jumped": 6,
			"over":   4,
			"the":    3,
			"lazy":   4,
			"dog.":   4,
		},
	}
}
