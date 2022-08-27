package game

import (
	"fmt"

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

	for i := 0; i < 0; i++ {
		d.Shuffle()
		res,_ := d.Deal()

		p1.Hand = res[0]
		p2.Hand = res[1]

		fmt.Printf("%d %d     %d %d \n", p1.Hand[0].Number, p1.Hand[0].Suit, p1.Hand[1].Number, p1.Hand[1].Suit)
		fmt.Printf("%d %d     %d %d \n", p2.Hand[0].Number, p2.Hand[0].Suit, p2.Hand[1].Number, p2.Hand[1].Suit)
	}

	evaluator.Evaluator([]player.Player{*p1,*p2}, [5]card.Card{card.Card{Number: 1, Suit: 3},card.Card{Number: 13, Suit: 2}, card.Card{Number: 8, Suit: 3}, card.Card{Number: 3, Suit: 2},card.Card{Number: 9, Suit: 0}})

}
