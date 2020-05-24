package models

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

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

func (deck *Deck) GetRemaining() int64 {
	return int64(len(deck.Cards))
}
