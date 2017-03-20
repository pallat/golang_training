package main

import "golang.org/x/tour/pic"
import "math/rand"

func Pic(dx, dy int) (data [][]uint8) {
	r := rand.New(rand.NewSource(2))

	row := []uint8{}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			row = append(row, uint8(r.Uint32()))
		}
		data = append(data, row)
	}
	return
}

func main() {
	pic.Show(Pic)
}
