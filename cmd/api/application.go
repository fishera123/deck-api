package api

import (
	"deck-api/pkg/models"
	"log"
)

type Application struct {
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
	DeckModel interface {
		Create(cards []string, isShuffled bool) models.Deck
		Get(id string) (models.Deck, error)
		Update(deck models.Deck)
	}
}
