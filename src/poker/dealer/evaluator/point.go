package evaluator

import (
	"poker/poker/playing_cards/card"
)

type Point struct {
	PlayerId int
	Role int
	UsedCards []card.Card
}

type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	return p[i].Role < p[j].Role
}
