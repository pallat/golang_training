package main

import (
	"fmt"
	"math"
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

	fmt.Println(math.Sqrt(g))
	fmt.Println(sqrt(g))
}

func sqrt(x float64) float64 {
	const (
		mask  = 0x7FF
		shift = 64 - 11 - 1
		bias  = 1023
	)
	// special cases
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	ix := math.Float64bits(x)
	// normalize x
	exp := int((ix >> shift) & mask)
	if exp == 0 { // subnormal x
		for ix&(1<<shift) == 0 {
			ix <<= 1
			exp--
		}
		exp++
	}
	exp -= bias // unbias exponent
	ix &^= mask << shift
	ix |= 1 << shift
	if exp&1 == 1 { // odd exp, double x to make it even
		ix <<= 1
	}
	exp >>= 1 // exp = exp/2, exponent of square root
	// generate sqrt(x) bit by bit
	ix <<= 1
	var q, s uint64               // q = sqrt(x)
	r := uint64(1 << (shift + 1)) // r = moving bit from MSB to LSB
	for r != 0 {
		t := s + r
		if t <= ix {
			s = t + r
			ix -= t
			q += r
		}
		ix <<= 1
		r >>= 1
	}
	// final rounding
	if ix != 0 { // remainder, result not exact
		q += q & 1 // round according to extra bit
	}
	ix = q>>1 + uint64(exp-1+bias)<<shift // significand + biased exponent
	return math.Float64frombits(ix)
}
