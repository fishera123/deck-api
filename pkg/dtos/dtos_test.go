package dto

import (
	"deck-api/pkg/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToCardsDto(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return CardsDto struct with correct data", func(t *testing.T) {
		deck := models.Deck{DeckId: uuid.New(), IsShuffled: false, Cards: []string{"10C"}}
		cardsDto := CardsDto{
			Cards: []CardDto{
				{
					Value: "10",
					Suit:  "CLUBS",
					Code:  "10C",
				},
			},
		}
		assert.Equal(ToCardsDto(deck.Cards), cardsDto)
	})
}

func TestToDeckDto(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return CardsDto struct with correct data", func(t *testing.T) {
		uuid := uuid.New()
		deck := models.Deck{DeckId: uuid, IsShuffled: false, Cards: []string{"10C"}}
		deckDto := CreatedDeckDto{
			DeckId:     uuid,
			IsShuffled: false,
			Remaining:  1,
		}
		assert.Equal(ToDeckDto(deck), deckDto)
	})
}

func TestToOpenDeckDto(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return CardsDto struct with correct data", func(t *testing.T) {
		uuid := uuid.New()
		deck := models.Deck{DeckId: uuid, IsShuffled: false, Cards: []string{"10C"}}
		deckDto := OpenDeckDto{
			CreatedDeckDto: CreatedDeckDto{
				DeckId:     uuid,
				IsShuffled: false,
				Remaining:  1,
			},
			Cards: []CardDto{
				{
					Value: "10",
					Suit:  "CLUBS",
					Code:  "10C",
				},
			},
		}
		assert.Equal(ToOpenDeckDto(deck), deckDto)
	})
}
