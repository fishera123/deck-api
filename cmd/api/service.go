package api

import (
	"github.com/fishera123/go-rest-api/pkg/models"
)

func (app *Application) CreateDeck(cardCodes string, isShuffled bool) (models.Deck, error) {
	cards := createDefaultCardSequence()
	return app.DeckModel.Create(cards, isShuffled), nil
}
