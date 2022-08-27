package dealer

import (
	"math/rand"
	"time"

	"poker/poker/playing_cards/card"
)

type Dealer struct {
	Players [2]bool
	Cards []card.Card
	Field [5]card.Card
	Pot int
	ActionHistory []ActionHistory
}

func NewDealer() *Dealer {
	dealer := Dealer{}
	dealer.init()
	return &dealer
}

func (dealer *Dealer) init() {
	dealer.Field = [5]card.Card{
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
	}
	dealer.Pot = 0
	dealer.ActionHistory = []ActionHistory{}

	for i := 0; i < card.SuitNum; i++ {
		for j := 1; j <= card.CardsNum; j++ {
			dealer.Cards = append(dealer.Cards, card.Card{Number: j, Suit: i})
		}
	}
}

func (dealer *Dealer) AddPlayer() int {
	for i, seated := range dealer.Players {
		if !seated {
			dealer.Players[i] = true
			return i
		}
	}
	return -1
}

func (dealer *Dealer) TakePlayer(playerId int) int {
	dealer.Players[playerId] = false
}

func (dealer *Dealer) Shuffle() {
	dealer.init()

	rand.Seed(time.Now().Unix())

	for i, _ := range dealer.Cards {
		index := rand.Intn(card.SuitNum * card.CardsNum)

		tmp := dealer.Cards[index]
		dealer.Cards[index] = dealer.Cards[i]
		dealer.Cards[i] = tmp
	}
}

func (dealer *Dealer) CalcPot(chip int) {
	dealer.Pot = dealer.Pot + chip
}
