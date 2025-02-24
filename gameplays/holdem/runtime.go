package holdem

import (
	"time"
)

type Game struct {
	settings *settings

	deck    *deck
	players *players
	button  uint8
	stage   string
	pool    uint64
	round   uint64
	logs    []*log
}

func NewGame(s *settings) *Game {
	g := &Game{
		settings: s,
		deck:     NewDeck(),
		button:   0,
		stage:    "",
		pool:     0,
		round:    0,
		logs:     make([]*log, 0),
	}

	g.writeLog(SystemUser, GameCreate)
	return g
}

func (g *Game) Deck() *deck {
	return g.deck
}

func (g *Game) Setting() *settings {
	return g.settings
}

func (g *Game) initButtonAndAnte() {

}

func (g *Game) writeLog(playerID string, action string) {
	l := &log{
		round:     g.round,
		player:    playerID,
		action:    action,
		timestamp: time.Now().UnixMilli(),
	}

	g.logs = append(g.logs, l)
}

//func (g *Game) Begin() (err error) {
//	g.round += g.round
//	g.writeLog(SystemUser, RoundBegin)
//
//	g.deck.Shuffle()
//	g.initButtonAndAnte()
//
//	blindNum := len(g.settings.blind)
//	for i, p := range g.players {
//
//		if g.button == i {
//
//		}
//
//		if p.hands, err = g.deck.DealCards(2); err != nil {
//			return fmt.Errorf("deal cards err : %w", err)
//		}
//	}
//	g.stage = PreFlop
//
//	return nil
//}
