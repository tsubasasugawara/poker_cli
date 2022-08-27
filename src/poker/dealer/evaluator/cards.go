package evaluator

import (
	"poker/poker/playing_cards/card"
)

// ハンドとボードを結合し、ソートするため
type Cards []card.Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cards) Less(i, j int) bool {
	return c[i].Number < c[j].Number
}
