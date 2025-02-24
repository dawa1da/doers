package holdem

import (
	playerBase "doers/resources/player"
	"doers/resources/poker"
)

type player struct {
	*playerBase.MetaData

	hands   []*poker.Card
	brought uint64
	chips   uint64
	next    *player
}

type players struct {
	head *player
}

func newPlayers() *players {
	return &players{}
}

func (ps *players) Add(p *player) {
	if ps.head == nil {
		ps.head = p
	} else {

	}
}
