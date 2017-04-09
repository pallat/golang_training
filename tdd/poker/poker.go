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
		return "Black wins - high card: " + rank(mxBlck)
	}

	if mxBlck == mxWhte {
		return "Tie"
	}

	return "White wins - high card: " + rank(mxWhte)
}

func rank(i int) string {
	mapping := map[int]string{
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "Jack",
		12: "Queen",
		13: "King",
		14: "Ace",
	}
	return mapping[i]
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
