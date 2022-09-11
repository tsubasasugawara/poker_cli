package evaluator

import (
	"poker/playing_cards/card"
)

type Cards []card.Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cards) Less(i, j int) bool {
	var numI, numJ int

	if c[i].Number == 0 {
		numI = 13
	} else {
		numI = c[i].Number
	}
	if c[j].Number == 0 {
		numJ = 13
	} else {
		numJ = c[j].Number
	}

	if numI == numJ {
		return c[i].Suit < c[j].Suit
	}
	return numI < numJ
}
