package mock

import (
	"deck-api/pkg/models"

	"github.com/google/uuid"
)

var deckMock = models.Deck{
	DeckId:     uuid.New(),
	IsShuffled: false,
	Cards:      []string{"10C"},
}

type DeckModel struct{}

func (model *DeckModel) Create(cards []string, isShuffled bool) models.Deck {
	return models.Deck{
		DeckId:     uuid.New(),
		IsShuffled: isShuffled,
		Cards:      cards,
	}
}

func (model *DeckModel) Get(id string) (models.Deck, error) {
	return deckMock, nil
}

func (model *DeckModel) Update(deck models.Deck) {

}
