package constants

import "strings"

const (
	ACE      = "A"
	TWO      = "2"
	THREE    = "3"
	FOUR     = "4"
	FIVE     = "5"
	SIX      = "6"
	SEVEN    = "7"
	EIGHT    = "8"
	NINE     = "9"
	TEN      = "10"
	JACK     = "J"
	QUEEN    = "Q"
	KING     = "K"
	SPADES   = "S"
	CLUBS    = "C"
	DIAMONDS = "D"
	HEARTS   = "H"
)

var SuitsSequence = []string{SPADES, DIAMONDS, CLUBS, HEARTS}
var Suits = map[string]string{
	SPADES:   "SPADES",
	CLUBS:    "CLUBS",
	DIAMONDS: "DIAMONDS",
	HEARTS:   "HEARTS",
}

var CardsSequence = []string{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
var Cards = map[string]string{
	ACE:   "ACE",
	JACK:  "JACK",
	QUEEN: "QUEEN",
	KING:  "KING",
}

var CardsString = strings.Join(CardsSequence, "")
var SuitsString = strings.Join(SuitsSequence, "")
