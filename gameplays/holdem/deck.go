package holdem

import (
	"doers/resources/poker"
)

type deck struct {
	*poker.DeckBase
}

func NewDeck() *deck {
	return &deck{
		DeckBase: poker.New(),
	}
}
