package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Deck struct {
	gorm.Model
	DeckId     uuid.UUID
	IsShuffled bool
	Cards      pq.StringArray `gorm:"type:varchar(255)[]"`
}
