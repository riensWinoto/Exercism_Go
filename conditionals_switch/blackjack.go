package main

import "fmt"

var (
	card                     string = "ace"
	card1, card2, dealerCard string = "two", "seven", "jack"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11

	case "two":
		return 2

	case "three":
		return 3

	case "four":
		return 4

	case "five":
		return 5

	case "six":
		return 6

	case "seven":
		return 7

	case "eight":
		return 8

	case "nine":
		return 9

	case "ten", "jack", "queen", "king":
		return 10

	default:
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var valueReturn string
	valueCard1 := ParseCard(card1)
	valueCard2 := ParseCard(card2)
	valueDealerCard := ParseCard(dealerCard)
	valueTotal := valueCard1 + valueCard2

	// P = Split; S = Stand; H = Hit; W = Automatically Win
	switch {
	case valueTotal <= 11:
		valueReturn = "H"

	case valueTotal >= 12 && valueTotal <= 16 && valueDealerCard >= 7:
		valueReturn = "H"

	case valueTotal >= 12 && valueTotal <= 16 && valueDealerCard < 7:
		valueReturn = "S"

	case valueTotal >= 17 && valueTotal <= 20:
		valueReturn = "S"

	case valueTotal == 21 && (valueDealerCard != 10 && valueDealerCard != 11):
		valueReturn = "W"

	case valueTotal == 21 && (valueDealerCard == 10 || valueDealerCard == 11):
		valueReturn = "S"

	case valueCard1 == 11 && valueCard2 == 11:
		valueReturn = "P"
	}
	return valueReturn
}

func main() {
	fmt.Println(ParseCard(card))
	fmt.Println(FirstTurn(card1, card2, dealerCard))
}
