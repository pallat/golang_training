package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	//var y int64
	var z float64

	for i := 1; i < 10; i++ {
		z = float64(i)
		res := z - ((z*z)-x)/(2*z)
		if ((res * res) - x) < 0.1 {
			k := res - 0.5
			for j := k; j < (res + 0.5); j += 0.1 {
				z = j
				res := z - ((z*z)-x)/(2*z)
				if ((res * res) - x) < 0.01 {
					k := res - 0.05
					for l := k; l < (res + 0.05); l += 0.01 {
						z = l
						res := z - ((z*z)-x)/(2*z)
						if ((res * res) - x) < 0.001 {
							return res
						}
					}
				}
			}
			//			z = float64(i-1)
			//			return z - ((z*z) - x) / (2 * z)
			return res
		}
	}
	return z * 1000
}

func main() {
	fmt.Println(Sqrt(3))
}
