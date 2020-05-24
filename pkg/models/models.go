package models

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

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

var ErrNoRecord = errors.New("models: no matching record found")

type Deck struct {
	gorm.Model
	DeckId     uuid.UUID
	IsShuffled bool
	Cards      pq.StringArray `gorm:"type:varchar(255)[]"`
}

func (deck *Deck) Shuffle() *Deck {
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return deck
}
