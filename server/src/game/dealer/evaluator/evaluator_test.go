package evaluator

import (
	"testing"
	"reflect"
	"sort"

	"poker/playing_cards/card"
)

type TestStructure struct {
	input Cards
	expectedCards Cards
	expectedRole int
}

func exec(t *testing.T, evl func(CardMap) (int, Cards), ts []TestStructure) {
	for i := 0; i < len(ts); i++ {
		cardMap := makeCardMap(ts[i].input)
		role, usedCards := evl(cardMap)
		if (role != ts[i].expectedRole) {
			t.Fatalf("Role is wrong. i = %d role = %d expectedRole = %d", i, role, ts[i].expectedRole)
		}

		sort.Sort(usedCards)
		sort.Sort(ts[i].expectedCards)
		if (!reflect.DeepEqual(usedCards, ts[i].expectedCards)) {
			t.Fatalf("Outputed cards are wrong. i = %d usedCards = %d expectedCards = %d",i, usedCards, ts[i].expectedCards)
		}
	}
}

func TestStraightFlash(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 2, Suit: 0},
				card.Card{Number: 3, Suit: 0},
				card.Card{Number: 4, Suit: 0},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 4, Suit: 0},
				card.Card{Number: 3, Suit: 0},
				card.Card{Number: 2, Suit: 0},
				card.Card{Number: 1, Suit: 0},
			},
			STRAIGHT_FLASH,
		},
		{
			Cards{
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 12, Suit: 0},
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 10, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 6, Suit: 0},
			},
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 12, Suit: 0},
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 10, Suit: 0},
				card.Card{Number: 9, Suit: 0},
			},
			STRAIGHT_FLASH,
		},
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 2, Suit: 0},
				card.Card{Number: 3, Suit: 0},
				card.Card{Number: 4, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 6, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards{},
			HIGH_CARD,
		},
	}

	exec(t, straightFlush, patterns)
}

func TestFourCard(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 3},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 1, Suit: 3},
				card.Card{Number: 9, Suit: 0},
			},
			FOUR_CARD,
		},
		{
			Cards{
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 6, Suit: 0},
			},
			Cards{},
			HIGH_CARD,
		},
	}

	exec(t, fourCard, patterns)
}

func TestFullHouse(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 5, Suit: 2},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 6, Suit: 1},
				card.Card{Number: 8, Suit: 2},
			},
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 5, Suit: 2},
				card.Card{Number: 5, Suit: 3},
			},
			FULL_HOUSE,
		},
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 5, Suit: 2},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 8, Suit: 1},
				card.Card{Number: 8, Suit: 2},
			},
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 8, Suit: 1},
				card.Card{Number: 8, Suit: 2},
			},
			FULL_HOUSE,
		},
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 9, Suit: 2},
				card.Card{Number: 5, Suit: 2},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 6, Suit: 1},
				card.Card{Number: 8, Suit: 2},
			},
			Cards{},
			HIGH_CARD,
		},
	}

	exec(t, fullHouse, patterns)
}

func TestStraight(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 2, Suit: 1},
				card.Card{Number: 3, Suit: 2},
				card.Card{Number: 4, Suit: 3},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 2, Suit: 1},
				card.Card{Number: 3, Suit: 2},
				card.Card{Number: 4, Suit: 3},
				card.Card{Number: 5, Suit: 0},
			},
			STRAIGHT,
		},
		{
			Cards{
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 10, Suit: 3},
				card.Card{Number: 9, Suit: 1},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 10, Suit: 3},
				card.Card{Number: 9, Suit: 0},
			},
			STRAIGHT,
		},
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 10, Suit: 1},
				card.Card{Number: 3, Suit: 2},
				card.Card{Number: 4, Suit: 3},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {},
			HIGH_CARD,
		},
	}

	exec(t, straight, patterns)
}

func TestFlash(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards{
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 4, Suit: 3},
				card.Card{Number: 3, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 3, Suit: 0},
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 7, Suit: 0},
				card.Card{Number: 9, Suit: 0},
			},
			FLASH,
		},
		{
			Cards{
				card.Card{Number: 11, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 0, Suit: 3},
				card.Card{Number: 4, Suit: 3},
				card.Card{Number: 3, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {},
			HIGH_CARD,
		},
	}

	exec(t, flash, patterns)
}

func TestThreeCards(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 4, Suit: 3},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			THREE_CARD,
		},
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 0, Suit: 3},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 0, Suit: 3},
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 7, Suit: 0},
				card.Card{Number: 1, Suit: 0},
			},
			THREE_CARD,
		},
		{
			Cards{
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 1, Suit: 2},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			Cards {
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 7, Suit: 0},
			},
			THREE_CARD,
		},
		{
			Cards{
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 1, Suit: 0},
				card.Card{Number: 1, Suit: 1},
				card.Card{Number: 5, Suit: 2},
				card.Card{Number: 9, Suit: 0},
				card.Card{Number: 5, Suit: 0},
				card.Card{Number: 6, Suit: 0},
			},
			Cards {},
			HIGH_CARD,
		},
	}

	exec(t, threeCard, patterns)
}

func TestTwoPair(t *testing.T) {
	patterns := []TestStructure {
		{
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 8, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10,Suit: 3},
			},
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 5, Suit: 1},
			},
			TWO_PAIR,
		},
		{
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 6, Suit: 3},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 6, Suit: 0},
				card.Card{Number: 5, Suit: 2},
				card.Card{Number: 10,Suit: 3},
			},
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 10, Suit: 3},
				card.Card{Number: 6, Suit: 3},
				card.Card{Number: 6, Suit: 0},
			},
			TWO_PAIR,
		},
		{
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 5, Suit: 3},
				card.Card{Number: 4, Suit: 1},
				card.Card{Number: 8, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10,Suit: 3},
			},
			Cards {},
			HIGH_CARD,
		},
	}

	exec(t, twoPair, patterns)
}

func TestOnePair(t *testing.T) {
	patterns := []TestStructure{
		{
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 4, Suit: 1},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 8, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10,Suit: 3},
			},
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 0, Suit: 2},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10, Suit: 3},
				card.Card{Number: 8, Suit: 0},
			},
			ONE_PAIR,
		},
		{
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 6, Suit: 2},
				card.Card{Number: 4, Suit: 1},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 8, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10,Suit: 3},
			},
			Cards {},
			HIGH_CARD,
		},
	}

	exec(t, onePair, patterns)
}

func TestHighCard(t *testing.T) {
	patterns := []TestStructure{
		{
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 6, Suit: 2},
				card.Card{Number: 4, Suit: 1},
				card.Card{Number: 5, Suit: 1},
				card.Card{Number: 8, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10,Suit: 3},
			},
			Cards {
				card.Card{Number: 0, Suit: 0},
				card.Card{Number: 6, Suit: 2},
				card.Card{Number: 8, Suit: 0},
				card.Card{Number: 12, Suit: 1},
				card.Card{Number: 10,Suit: 3},
			},
			HIGH_CARD,
		},
	}

	exec(t, highCard, patterns)
}
