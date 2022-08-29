package evaluator

import (
	"testing"

	"poker/poker/playing_cards/card"
)

func TestStraightFlash(t *testing.T) {
	input := [][]card.Card {
		[]card.Card{
			card.Card{Number: 1, Suit: 0},
			card.Card{Number: 2, Suit: 0},
			card.Card{Number: 3, Suit: 0},
			card.Card{Number: 4, Suit: 0},
			card.Card{Number: 5, Suit: 0},
			card.Card{Number: 9, Suit: 0},
			card.Card{Number: 7, Suit: 0},
		},
		[]card.Card{
			card.Card{Number: 0, Suit: 0},
			card.Card{Number: 12, Suit: 0},
			card.Card{Number: 11, Suit: 0},
			card.Card{Number: 10, Suit: 0},
			card.Card{Number: 9, Suit: 0},
			card.Card{Number: 5, Suit: 0},
			card.Card{Number: 6, Suit: 0},
		},
		[]card.Card{
			card.Card{Number: 1, Suit: 0},
			card.Card{Number: 2, Suit: 0},
			card.Card{Number: 3, Suit: 0},
			card.Card{Number: 4, Suit: 0},
			card.Card{Number: 9, Suit: 0},
			card.Card{Number: 6, Suit: 0},
			card.Card{Number: 7, Suit: 0},
		},
	}
	for i := 0; i < len(input); i++ {
		cardMap := MakeCardMap(input[i])

		role, usedCards := StraightFlush(cardMap)

		if (role != STRAIGHT_FLASH && len(usedCards) == 5) || (role == STRAIGHT_FLASH && len(usedCards) != 5) {
			t.Fatalf("role = %d  usedCards = {%d %d}, {%d %d}, {%d %d}, {%d %d}, {%d %d}", role, usedCards[0].Number, usedCards[0].Suit, usedCards[1].Number, usedCards[1].Suit, usedCards[2].Number, usedCards[2].Suit, usedCards[3].Number, usedCards[3].Suit, usedCards[4].Number, usedCards[4].Suit)
		}
	}
}

func TestFourCard(t *testing.T) {
	input := [][]card.Card {
		[]card.Card{
			card.Card{Number: 1, Suit: 0},
			card.Card{Number: 1, Suit: 1},
			card.Card{Number: 1, Suit: 2},
			card.Card{Number: 1, Suit: 3},
			card.Card{Number: 5, Suit: 0},
			card.Card{Number: 9, Suit: 0},
			card.Card{Number: 7, Suit: 0},
		},
		[]card.Card{
			card.Card{Number: 0, Suit: 0},
			card.Card{Number: 1, Suit: 0},
			card.Card{Number: 1, Suit: 1},
			card.Card{Number: 1, Suit: 2},
			card.Card{Number: 9, Suit: 0},
			card.Card{Number: 5, Suit: 0},
			card.Card{Number: 6, Suit: 0},
		},
	}
	for i := 0; i < len(input); i++ {
		cardMap := MakeCardMap(input[i])

		role, usedCards := FourCard(cardMap)

		if (role != FOUR_CARD && len(usedCards) == 5) || (role == FOUR_CARD && len(usedCards) != 5) {
			t.Fatalf("role = %d  usedCards = {%d %d}, {%d %d}, {%d %d}, {%d %d}, {%d %d}", role, usedCards[0].Number, usedCards[0].Suit, usedCards[1].Number, usedCards[1].Suit, usedCards[2].Number, usedCards[2].Suit, usedCards[3].Number, usedCards[3].Suit, usedCards[4].Number, usedCards[4].Suit)
		}
	}
}
