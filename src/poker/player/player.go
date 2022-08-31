package player

import (
	"poker/poker/playing_cards/card"
	"poker/poker/dealer"
)

type Player struct {
	Id int
	Hand [2]card.Card
	Stack int
}

func NewPlayer(stack int) *Player {
	player := Player{Stack: stack}
	return &player
}

func (p *Player) Bet(chip int) (int, int) {
	if chip > p.Stack {
		return dealer.ERROR, 0
	} else if chip == p.Stack {
		p.Stack = 0
		return dealer.ALLIN, chip
	} else if chip > 0 {
		p.Stack = p.Stack - chip
		return dealer.BET, chip
	} else {
		return dealer.CHECK, chip
	}
}

func (p *Player) Win(chip int) {
	p.Stack = p.Stack + chip
}
