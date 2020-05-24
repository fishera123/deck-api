package api

import (
	"deck-api/pkg/models"
	"errors"
	"fmt"
	"strings"
)

func (app *Application) CreateDeck(cardCodes string, isShuffled bool) (models.Deck, error) {
	var cards []string
	var err error
	if len(cardCodes) > 0 {
		cards, err = createRangeCardSequence(cardCodes)
	} else {
		cards = createDefaultCardSequence()
	}
	emptyDeck := models.Deck{}
	if err != nil {
		return emptyDeck, err
	}

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

func createRangeCardSequence(cardCodeRange string) (codes []string, err error) {
	codes = strings.Split(cardCodeRange, ",")
	if isValidCardCode(codes) {
		return codes, nil
	}
	return codes, errors.New("invalid card codes")
}
