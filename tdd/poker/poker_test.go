package poker

import "testing"

func TestWhiteWinsByRank(t *testing.T) {
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

func TestBlackWinsByRank(t *testing.T) {
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

func TestBlackWinsByHighRankIsKing(t *testing.T) {
	white := []card{
		card{rank: "2", suit: "H"},
		card{rank: "3", suit: "D"},
		card{rank: "5", suit: "S"},
		card{rank: "9", suit: "C"},
		card{rank: "Q", suit: "D"},
	}
	black := []card{
		card{rank: "2", suit: "C"},
		card{rank: "3", suit: "H"},
		card{rank: "4", suit: "S"},
		card{rank: "8", suit: "C"},
		card{rank: "K", suit: "H"},
	}

	result := compare(black, white)
	if result != "Black wins - high card: King" {
		t.Error("result should be 'Black wins - high card: King' but was", result)
	}
}

func TestTie(t *testing.T) {
	black := []card{
		card{rank: "2", suit: "H"},
		card{rank: "3", suit: "D"},
		card{rank: "5", suit: "S"},
		card{rank: "9", suit: "C"},
		card{rank: "K", suit: "D"},
	}
	white := []card{
		card{rank: "2", suit: "D"},
		card{rank: "3", suit: "H"},
		card{rank: "5", suit: "C"},
		card{rank: "9", suit: "S"},
		card{rank: "K", suit: "H"},
	}

	result := compare(black, white)
	if result != "Tie" {
		t.Error("result should be 'Tie' but was", result)
	}
}

func TestBlackWinsByAPairOfTwo(t *testing.T) {
	black := []card{
		card{rank: "2", suit: "H"},
		card{rank: "2", suit: "D"},
		card{rank: "5", suit: "S"},
		card{rank: "9", suit: "C"},
		card{rank: "K", suit: "D"},
	}
	white := []card{
		card{rank: "3", suit: "H"},
		card{rank: "4", suit: "D"},
		card{rank: "5", suit: "C"},
		card{rank: "9", suit: "S"},
		card{rank: "K", suit: "H"},
	}

	result := compare(black, white)
	if result != "Black Wins - Pair" {
		t.Error("result should be 'Black Wins - Pair' but was", result)
	}
}

func TestWhiteWinsByATwoPairs(t *testing.T) {
	black := []card{
		card{rank: "2", suit: "H"},
		card{rank: "2", suit: "D"},
		card{rank: "5", suit: "S"},
		card{rank: "9", suit: "C"},
		card{rank: "K", suit: "D"},
	}
	white := []card{
		card{rank: "3", suit: "H"},
		card{rank: "3", suit: "D"},
		card{rank: "8", suit: "C"},
		card{rank: "8", suit: "S"},
		card{rank: "K", suit: "H"},
	}

	result := compare(black, white)
	if result != "White Wins - Pairs" {
		t.Error("result should be 'White Wins - Pairs' but was", result)
	}
}
