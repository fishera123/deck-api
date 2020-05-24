package api

import (
	mock "deck-api/pkg/database/mocks"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestApplication(t *testing.T) *Application {
	// mock out application dependencies
	return &application{
		errorLog:  log.New(ioutil.Discard, "", 0),
		infoLog:   log.New(ioutil.Discard, "", 0),
		deckModel: &mock.DeckModel{}, //pass mock model which implements deck model interface
	}
}

func TestCreateDeck(t *testing.T) {
	assert := assert.New(t)
	app := newTestApplication(t)
	t.Run("should create deck with default sequence when request card codes are empty", func(t *testing.T) {
		cardCodes := ""
		isShuffled := false
		deck, _ := app.CreateDeck(cardCodes, isShuffled)
		assert.Len(deck.Cards, 52)
	})
	t.Run("should create deck with partial cards when request card codes exist", func(t *testing.T) {
		cardCodes := "9C,10C"
		isShuffled := false
		deck, _ := app.CreateDeck(cardCodes, isShuffled)
		assert.Len(deck.Cards, 2)
	})
}
