package poker

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

const (
	CardSuitSpade   = "♠️"
	CardSuitHeart   = "♥️"
	CardSuitClub    = "♣️"
	CardSuitDiamond = "♦️"
	CardSuitRed     = "red️"
	CardSuitBlack   = "black"
)

var defaultCardSuits = []string{
	CardSuitSpade,
	CardSuitHeart,
	CardSuitClub,
	CardSuitDiamond,
}

func DefaultCardSuits() []string {
	bak := make([]string, len(defaultCardSuits))
	copy(bak, defaultCardSuits)
	return bak
}

var defaultCardExtraSuits = []string{
	CardSuitRed,
	CardSuitBlack,
}

func DefaultCardExtraSuits() []string {
	bak := make([]string, len(defaultCardExtraSuits))
	copy(bak, defaultCardExtraSuits)
	return bak
}

const (
	CardPoint2 uint8 = iota + 2
	CardPoint3
	CardPoint4
	CardPoint5
	CardPoint6
	CardPoint7
	CardPoint8
	CardPoint9
	CardPoint10
	CardPointJ
	CardPointQ
	CardPointK
	CardPointA
	CardPointJoker
)

const (
	StringJ     = "J"
	StringQ     = "Q"
	StringK     = "K"
	StringA     = "A"
	StringJoker = "Joker"
)

var defaultCardPointsRank = []uint8{
	CardPoint2,
	CardPoint3,
	CardPoint4,
	CardPoint5,
	CardPoint6,
	CardPoint7,
	CardPoint8,
	CardPoint9,
	CardPoint10,
	CardPointJ,
	CardPointQ,
	CardPointK,
	CardPointA,
}

func DefaultCardPointsRank() []uint8 {
	bak := make([]uint8, len(defaultCardPointsRank))
	copy(bak, defaultCardPointsRank)
	return bak
}

type Card struct {
	suit  string
	point uint8
}

func (c *Card) Suit() string {
	return c.suit
}

func (c *Card) Point() uint8 {
	return c.point
}

func (c *Card) String() string {
	var str string

	switch c.point {
	case CardPointJ:
		str = StringJ
	case CardPointQ:
		str = StringQ
	case CardPointK:
		str = StringK
	case CardPointA:
		str = StringA
	case CardPointJoker:
		str = StringJoker
	default:
		str = strconv.Itoa(int(c.point))
	}

	return fmt.Sprintf("%s%s", c.suit, str)
}

func IsSuitAndPointLegal(suit string, point uint8) bool {
	if point == CardPointJoker && slices.Contains(defaultCardExtraSuits, suit) {
		return true
	} else if slices.Contains(defaultCardSuits, suit) && slices.Contains(defaultCardPointsRank, point) {
		return true
	} else {
		return false
	}
}
