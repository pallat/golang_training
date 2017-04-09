package poker

import "testing"

func TestWhiteWinsByScore(t *testing.T) {
	black := []card{
		card{rank: "2", suit: "H"},
		card{rank: "3", suit: "D"},
		card{rank: "5", suit: "S"},
		card{rank: "9", suit: "C"},
		card{rank: "K", suit: "D"},
	}
	white := []card{
		card{rank: "2", suit: "C"},
		card{rank: "3", suit: "H"},
		card{rank: "4", suit: "S"},
		card{rank: "8", suit: "C"},
		card{rank: "A", suit: "H"},
	}

	result := compare(black, white)
	if result != "White wins - high card: Ace" {
		t.Error("result should be 'White wins - high card: Ace' but was", result)
	}
}

func TestBlackWinsByScore(t *testing.T) {
	white := []card{
		card{rank: "2", suit: "H"},
		card{rank: "3", suit: "D"},
		card{rank: "5", suit: "S"},
		card{rank: "9", suit: "C"},
		card{rank: "K", suit: "D"},
	}
	black := []card{
		card{rank: "2", suit: "C"},
		card{rank: "3", suit: "H"},
		card{rank: "4", suit: "S"},
		card{rank: "8", suit: "C"},
		card{rank: "A", suit: "H"},
	}

	result := compare(black, white)
	if result != "Black wins - high card: Ace" {
		t.Error("result should be 'Black wins - high card: Ace' but was", result)
	}
}
