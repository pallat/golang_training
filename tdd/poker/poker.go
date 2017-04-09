package poker

import (
	"strconv"
)

type card struct {
	rank string
	suit string
}

type cards []card

func compare(black, white cards) string {
	mxBlck, mxWhte := 0, 0

	for _, v := range black {
		if val(v.rank) > mxBlck {
			mxBlck = val(v.rank)
		}
	}
	for _, v := range white {
		if val(v.rank) > mxWhte {
			mxWhte = val(v.rank)
		}
	}

	if mxBlck > mxWhte {
		return "Black wins - high card: Ace"
	}

	return "White wins - high card: Ace"
}

func val(s string) int {
	mapping := map[string]string{
		"T": "10",
		"J": "11",
		"Q": "12",
		"K": "13",
		"A": "14",
	}

	if v, ok := mapping[s]; ok {
		s = v
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return n
}
