package dto

import (
	"deck-api/pkg/constants"
	"deck-api/pkg/models"

	"github.com/google/uuid"
)

type CreatedDeckDto struct {
	DeckId     uuid.UUID `json:"deck_id"`
	IsShuffled bool      `json:"shuffled"`
	Remaining  int64     `json:"remaining"`
}

type OpenDeckDto struct {
	CreatedDeckDto
	Cards []CardDto `json:"cards"`
}

type CardDto struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type CardsDto struct {
	Cards []CardDto `json:"cards"`
}

func getCardValue(key string) string {
	value, exists := constants.Cards[key]
	if exists {
		return value
	} else {
		return key
	}
}

func convertCardToDto(card string) CardDto {
	byteSlice := []byte(card)
	suitValue := string(byteSlice[len(byteSlice)-1])
	cardValue := string(byteSlice[:len(byteSlice)-1])

	return CardDto{Value: getCardValue(cardValue), Suit: constants.Suits[suitValue], Code: card}
}

func getCardsDto(cards []string) []CardDto {
	cardsDto := []CardDto{}
	for _, card := range cards {
		cardsDto = append(cardsDto, convertCardToDto(card))
	}

	return cardsDto
}

func ToCardsDto(cards []string) CardsDto {
	return CardsDto{Cards: getCardsDto(cards)}
}

func ToDeckDto(deck models.Deck) CreatedDeckDto {
	return CreatedDeckDto{DeckId: deck.DeckId, Remaining: deck.GetRemaining(), IsShuffled: deck.IsShuffled}
}

func ToOpenDeckDto(deck models.Deck) OpenDeckDto {
	return OpenDeckDto{ToDeckDto(deck), getCardsDto(deck.Cards)}
}
