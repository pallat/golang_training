package main

import (
	"fmt"
	"os"
)

func Sqrt(x float64) float64 {
	var z float64

	for i := 1; i < 10; i++ {
		z = float64(i)
		res := fomular(z, x)
		if isclosest(res, 1, x) {
			for j := 10.00; j < 10000000000; j *= 10 {
				res = closest(res, 1/j, x)
			}
			return res
		}
	}
	return x
}

func closest(near, ratio, x float64) float64 {
	min := near - (ratio * 5)
	max := near + (ratio * 5)
	for i := min; i <= max; i += ratio {
		res := fomular(i, x)
		if (res * res) == x {
			return res
		}
		if isclosest(res, ratio, x) {
			return res
		}
	}
	return 0
}

func isclosest(res, ratio, x float64) bool {
	return ((res * res) - x) < (ratio / 3)
}

func fomular(z, x float64) float64 {
	return z - (((z * z) - x) / (2 * z))
}

func main() {
	var g float64
	fmt.Fscanln(os.Stdin, &g)
	fmt.Println(Sqrt(g))
}
