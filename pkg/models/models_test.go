package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	assert := assert.New(t)

	t.Run("should shuffle cards slice correctly", func(t *testing.T) {
		cards := []string{"9C", "10C"}
		deck := Deck{DeckId: uuid.New(), IsShuffled: false, Cards: cards}
		assert.NotEqual(deck.Shuffle().Cards, cards)
	})
}
