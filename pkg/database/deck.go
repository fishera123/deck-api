package db

import (
	"deck-api/pkg/models"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// wrap DB connection pool in Deck model, to be injected wherever we initiate db
type DeckModel struct {
	DB *gorm.DB
}

func (model *DeckModel) Create(cards []string, isShuffled bool) models.Deck {
	deck := models.Deck{DeckId: uuid.New(), Cards: cards, IsShuffled: isShuffled}
	if deck.IsShuffled {
		deck.Shuffle()
	}
	model.DB.NewRecord(deck)
	model.DB.Create(&deck)

	return deck
}

func (model *DeckModel) Get(id string) (models.Deck, error) {
	var deck = models.Deck{}
	if model.DB.First(&deck, "deck_id = ?", id).RecordNotFound() {
		return deck, models.ErrNoRecord
	}
	return deck, nil
}

func (model *DeckModel) Update(deck models.Deck) {
	model.DB.Save(&deck)
}
