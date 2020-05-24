package api

import (
	"deck-api/pkg/constants"
	"deck-api/pkg/models"
	"errors"
	"fmt"
	"strings"
)

// validation should be handled in middleware, would be more cleaner
// todo: this method can return an array of card codes which failed validation and return more detailed errors in response.
func isValidCardCode(cards []string) bool {
	for _, card := range cards {
		// todo: remove code duplication
		byteSlice := []byte(card)
		suitValue := string(byteSlice[len(byteSlice)-1])
		cardValue := string(byteSlice[:len(byteSlice)-1])

		if !strings.Contains(constants.CardsString, cardValue) || !strings.Contains(constants.SuitsString, suitValue) {
			return false
		}
	}
	return true
}

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

func (app *Application) DrawDeck(deck models.Deck, count int64) (drawnCards []string, err error) {
	remainingCards := deck.GetRemaining()
	if remainingCards == 0 {
		return drawnCards, errors.New("no remain cards left")
	}
	if remainingCards < count {
		count = remainingCards
	}
	drawnCards = deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	app.DeckModel.Update(deck)
	return drawnCards, nil
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
