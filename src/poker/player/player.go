package player

import (
	"poker/poker/dealer"
)

type Player struct {
	Identification Number int
	Hands [2]int
	Stack int
}

func NewPlayer(stack int) *Player {
	player := Player{Hands: []int{0,0}, Stack: stack}
	return &player
}

// bet処理
func (p *Player) Bet(chip int) int {
	if chip > p.Stack {
		return dealer.ERROR
	} else if chip == p.Stack {
		p.Stack = 0
		return dealer.ALLIN
	} else if chip > 0 {
		p.Stack = p.Stack - chip
		return dealer.BET
	} else {
		return dealer.CHECK
	}
}

func (p *Player) Win(chip int) {
	p.Stack = p.Stack + chip
}
