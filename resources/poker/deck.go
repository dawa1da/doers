package poker

import (
	"fmt"
	"math/rand"
)

var defaultDeckCards = func() []*Card {
	cards := make([]*Card, 0, 52)

	for _, suit := range defaultCardSuits {
		for _, point := range defaultCardPointsRank {
			desk := &Card{
				suit:  suit,
				point: point,
			}
			cards = append(cards, desk)
		}
	}

	return cards
}()

type DeckBase struct {
	cards  []*Card
	cursor uint8
}

type DeckFactory interface {
	AddCard(suit string, point uint8) (err error)
	RemoveCard(index uint8) (err error)
	DealCards(num uint8) (card []*Card, err error)
	Shuffle()
}

func New() *DeckBase {
	deck := &DeckBase{
		cards:  make([]*Card, len(defaultDeckCards)),
		cursor: 0,
	}
	copy(deck.cards, defaultDeckCards)

	return deck
}

func (d *DeckBase) AddCard(suit string, point uint8) (err error) {
	if !IsSuitAndPointLegal(suit, point) {
		return fmt.Errorf("card is illegality")
	}

	d.cards = append(d.cards, &Card{
		suit:  suit,
		point: point,
	})

	return nil
}

func (d *DeckBase) RemoveCard(index uint8) (err error) {
	if d.cursor > index {
		return fmt.Errorf("card is out of range")
	}

	d.cards = append(d.cards[:index], d.cards[index+1:]...)

	return nil
}

func (d *DeckBase) DealCards(num uint8) (card []*Card, err error) {
	if d.cursor+num > uint8(len(d.cards)) {
		return nil, fmt.Errorf("deck index out of range")
	}

	defer func() { d.cursor += num }()

	return d.cards[d.cursor : d.cursor+num], nil
}

func (d *DeckBase) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	d.cursor = 0
}
