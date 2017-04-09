package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
	}

	m["Bell Labs"] = Vertex{40.68444, -74.39967}

	fmt.Println(m["Bell Labs"].Lat)

	fmt.Println(m["Bell Labs"])
}