package game

import (
	"poker/poker/dealer"
	"poker/poker/player"

	"poker/poker/playing_cards/card"
	"poker/poker/dealer/evaluator"
)

func Start() {
	d := dealer.NewDealer()

	p1 := player.NewPlayer(1000)
	p2 := player.NewPlayer(1000)

	p1id, err1 := d.AddPlayer()
	if err1 != nil {
		return
	}
	p1.Id = p1id

	p2id,err2 := d.AddPlayer()
	if err2 != nil {
		return
	}
	p2.Id = p2id

	for i := 0; i < 10; i++ {
		d.Shuffle()
		res,_ := d.Deal()

		p1.Hand = res[0]
		p2.Hand = res[1]

		evaluator.Evaluator(
		[]player.Player{*p1,*p2},
		[5]card.Card{
			card.Card{Number: 1, Suit: 0},
			card.Card{Number: 13, Suit: 0},
			card.Card{Number: 12, Suit: 0},
			card.Card{Number: 10, Suit: 0},
			card.Card{Number: 11, Suit: 0}},
		)
	}
}
