package holdem

const SystemUser = "system"

const (
	GameCreate = "create"
	GameDelete = "delete"

	RoundBegin     = "begin"
	RoundOver      = "over"
	RoundDealCards = "dealCards"

	ActionCheck = "check"
	ActionBet   = "bet"
	ActionCall  = "call"
	ActionRaise = "raise"
	ActionFold  = "fold"
	ActionAllin = "allin"
)

const (
	SB = "SB"
	BB = "BB"

	UTG  = "UTG"
	UTG1 = "UTG+1"
	UTG2 = "UTG+2"

	LJ  = "LJ"
	HJ  = "HJ"
	CO  = "CO"
	BTN = "BTN"
)

const (
	PreFlop = "pre-flop"
	Flop    = "flop"
	Turn    = "turn"
	River   = "river"
)

const (
	HighCard      = "high card"
	OnePair       = "one pair"
	TwoPair       = "two pair"
	ThreeOfAKind  = "three of a kind"
	Straight      = "straight"
	Flush         = "flush"
	FullHouse     = "full house"
	FourOfAKind   = "four of a kind"
	StraightFlush = "straight flush"
	RoyalFlush    = "royal flush"
)

var defaultShapeRankMap = map[string]uint8{
	HighCard:      1,
	OnePair:       2,
	TwoPair:       3,
	ThreeOfAKind:  4,
	Straight:      5,
	Flush:         6,
	FullHouse:     7,
	FourOfAKind:   8,
	StraightFlush: 9,
	RoyalFlush:    10,
}
