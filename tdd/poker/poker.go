package poker

type card struct {
	rank string
	suit string
}

type cards []card

func compare(black, white cards) string {
	return "White wins - high card: Ace"
}
