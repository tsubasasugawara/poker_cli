package main

import (
	"fmt"
	"poker/poker/playing_cards/card"
	"poker/poker/dealer"
)

func main() {
	fmt.Println("Hello World")
	d := dealer.NewDealer()
	d.Shuffle()

	for i := 0; i < card.SuitNum * card.CardsNum; i++ {
		fmt.Printf("%d, %d\n", d.Cards[i].Number, d.Cards[i].Suit)
	}
}
