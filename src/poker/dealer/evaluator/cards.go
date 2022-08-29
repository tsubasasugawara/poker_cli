package evaluator

import (
	"poker/poker/playing_cards/card"
)

type Cards []card.Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cards) Less(i, j int) bool {
	if c[i].Number == c[j].Number {
		return c[i].Suit < c[j].Suit
	}
	return c[i].Number < c[j].Number
}
