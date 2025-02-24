package holdem

import (
	"errors"
	"fmt"
	"sort"

	"gonum.org/v1/gonum/stat/combin"

	"doers/resources/poker"
)

var (
	equalityErr = errors.New("card point is equal")
)

type fiveCard [5]*poker.Card
type fives []*fiveCard

func CombineFives(cards []*poker.Card) (fs fives, err error) {
	if len(cards) < 5 {
		return nil, fmt.Errorf("length of cards must be >= 5")
	}

	list := combin.Combinations(len(cards), 5)
	fs = make(fives, 0, len(list))
	for _, data := range list {
		fc := new(fiveCard)
		for i, datum := range data {
			fc[i] = cards[datum]
		}
		fs = append(fs, fc)
	}
	return fs, nil
}

func (c *fiveCard) isFlush() bool {
	firstSuit := c[0].Suit()
	for _, card := range c[1:] {
		if card.Suit() != firstSuit {
			return false
		}
	}
	return true
}

func (c *fiveCard) isStraight() bool {
	c.sort()

	// A5432 特殊判断
	if c[1].Point() == poker.CardPoint5 && c[0].Point() == poker.CardPointA {
		return true
	}

	// 判断是否连续
	for i := 1; i < len(c); i++ {
		if c[i-1].Point() != c[i].Point()+1 {
			return false
		}
	}
	return true
}

func (c *fiveCard) isStraightFlush() bool {
	return c.isFlush() && c.isStraight()
}

func (c *fiveCard) isRoyalFlush() bool {
	return c.isFlush() && c.isStraight() && c[0].Point() == poker.CardPointA
}

func (c *fiveCard) Shape() string {
	c.sort()

	pointMap := make(map[uint8]uint8)
	for _, card := range c {
		pointMap[card.Point()]++
	}

	switch len(pointMap) {
	case 2:
		//  {a,a,a,b,b} || {a,a,a,a,b}
		for _, count := range pointMap {
			if count == 2 || count == 3 {
				return FullHouse
			} else {
				return FourOfAKind
			}
		}
	case 3:
		// {a,a,a,b,c} || {a,a,b,b,c}
		for _, count := range pointMap {
			if count == 2 {
				return TwoPair
			} else if count == 3 {
				return ThreeOfAKind
			} else {
				continue
			}
		}
	case 4:
		// {a,a,b,c,d}
		return OnePair
	case 5:
		// {a,b,c,d,e}
		if c.isStraight() {
			if c.isFlush() {
				if c[4].Point() == poker.CardPointA {
					return RoyalFlush
				}
				return StraightFlush
			}
			return Straight
		} else if c.isFlush() {
			if c.isStraight() {
				if c[4].Point() == poker.CardPointA {
					return RoyalFlush
				}
				return StraightFlush
			}
			return Flush
		} else {
			return HighCard
		}
	}

	return ""
}

func (c *fiveCard) sort() {
	slice := c[:]

	//按大小排序
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Point() > slice[j].Point()
	})

	countMap := make(map[uint8]uint8)
	for _, card := range slice {
		countMap[card.Point()]++
	}

	// 按数量排序
	sort.Slice(slice, func(i, j int) bool {
		return countMap[slice[i].Point()] > countMap[slice[j].Point()]
	})
}

func (f fives) sort() {
	sort.Slice(f, func(i, j int) bool {
		if f[i].Shape() == f[j].Shape() {
			for n := 0; n < 5; n++ {
				if f[i][n] == f[j][n] {
					continue
				}
				return f[i][n].Point() > f[j][n].Point()
			}
		}

		return defaultShapeRankMap[f[i].Shape()] > defaultShapeRankMap[f[j].Shape()]
	})
}

func (f fives) MaxOnes() fives {
	f.sort()

	ones := make(fives, 0)
	ones = append(ones, f[0])
	max := f[0]

	for i := 1; i < len(f); i++ {
		if max.Shape() == f[i].Shape() {
			for n := 0; n < 5; n++ {
				if max[n].Point() != f[i][n].Point() {
					break
				} else {
					if n == 4 {
						ones = append(ones, f[i])
					}
				}
			}
		} else {
			break
		}
	}

	return ones
}
