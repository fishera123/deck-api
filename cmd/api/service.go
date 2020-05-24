package api

import (
	"deck-api/pkg/models"
	"fmt"
)

func (app *Application) CreateDeck(cardCodes string, isShuffled bool) (models.Deck, error) {
	cards := createDefaultCardSequence()
	return app.DeckModel.Create(cards, isShuffled), nil
}

func createDefaultCardSequence() (codes []string) {
	for _, suit := range models.SuitsSequence {
		for _, card := range models.CardsSequence {
			code := fmt.Sprintf("%s%s", card, suit)
			codes = append(codes, code)
		}
	}

	return codes
}
