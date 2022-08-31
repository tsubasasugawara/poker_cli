package judge

import (
	"testing"
	"reflect"
	"sort"

	"poker/poker/dealer/evaluator"
	"poker/poker/playing_cards/card"
)

type TestStructure struct {
	input evaluator.Roles
	expectedWinner Ids
}

func exec(t *testing.T, evl func(candidates evaluator.Roles) (Ids), ts []TestStructure) {
	for i, v := range ts {
		res := evl(v.input)

		sort.Sort(res)
		sort.Sort(v.expectedWinner)

		if !reflect.DeepEqual(res, v.expectedWinner) {
			t.Fatalf("Outputed ids are wrong. i = %d res = %d expectedWinner = %d",i, res, v.expectedWinner)
		}
	}
}

func TestStraightFlash(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT_FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 4, Suit: 0},
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT_FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 6, Suit: 0},
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 4, Suit: 0},
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 2, Suit: 0},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT_FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 0},
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT_FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 0},
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 0},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT_FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 0},
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT_FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 0},
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 0},
					},
				},
			},
			[]int {0,3},
		},
	}

	exec(t, straightFlush, patterns)
}
func TestFourCard(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 0, Suit: 1},
						card.Card{Number: 0, Suit: 2},
						card.Card{Number: 0, Suit: 3},
						card.Card{Number: 1, Suit: 3},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 1, Suit: 1},
						card.Card{Number: 1, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 2, Suit: 0},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 1, Suit: 3},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 1, Suit: 1},
						card.Card{Number: 1, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 2, Suit: 0},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 12, Suit: 1},
						card.Card{Number: 12, Suit: 2},
						card.Card{Number: 12, Suit: 3},
						card.Card{Number: 9, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 12, Suit: 1},
						card.Card{Number: 12, Suit: 2},
						card.Card{Number: 12, Suit: 3},
						card.Card{Number: 8, Suit: 0},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
						card.Card{Number: 2, Suit: 2},
						card.Card{Number: 2, Suit: 3},
						card.Card{Number: 1, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 1, Suit: 1},
						card.Card{Number: 1, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 12, Suit: 0},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 12, Suit: 1},
						card.Card{Number: 12, Suit: 2},
						card.Card{Number: 12, Suit: 3},
						card.Card{Number: 8, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FOUR_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 12, Suit: 1},
						card.Card{Number: 12, Suit: 2},
						card.Card{Number: 12, Suit: 3},
						card.Card{Number: 8, Suit: 0},
					},
				},
			},
			[]int {0,3},
		},
	}

	exec(t, fourCard, patterns)
}

func TestFullHouse(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 5, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
					},
				},
			},
			[]int {0, 3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 1, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 2, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 2, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FULL_HOUSE,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 1, Suit: 1},
						card.Card{Number: 1, Suit: 2},
						card.Card{Number: 13, Suit: 0},
						card.Card{Number: 13, Suit: 1},
					},
				},
			},
			[]int {0},
		},
	}

	exec(t, fullHouse, patterns)
}

func TestFlush(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 0},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 0},
					},
				},
			},
			[]int {0, 3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 13, Suit: 0},
						card.Card{Number: 4, Suit: 0},
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.FLUSH,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 1},
						card.Card{Number: 11, Suit: 1},
						card.Card{Number: 10, Suit: 1},
						card.Card{Number: 9, Suit: 1},
						card.Card{Number: 7, Suit: 1},
					},
				},
			},
			[]int {0},
		},
	}

	exec(t, flush, patterns)
}

func TestStraight(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 10, Suit: 1},
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 7, Suit: 2},
						card.Card{Number: 6, Suit: 0},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 7, Suit: 0},
						card.Card{Number: 6, Suit: 3},
					},
				},
			},
			[]int {0, 3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 7, Suit: 2},
						card.Card{Number: 6, Suit: 0},
						card.Card{Number: 5, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 2},
						card.Card{Number: 7, Suit: 0},
						card.Card{Number: 6, Suit: 3},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 12, Suit: 3},
						card.Card{Number: 11, Suit: 2},
						card.Card{Number: 10, Suit: 0},
						card.Card{Number: 9, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 0},
						card.Card{Number: 10, Suit: 2},
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 3},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 4, Suit: 3},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 1, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.STRAIGHT,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 0},
						card.Card{Number: 10, Suit: 2},
						card.Card{Number: 9, Suit: 0},
						card.Card{Number: 8, Suit: 3},
					},
				},
			},
			[]int {3},
		},
	}

	exec(t, straight, patterns)
}

func TestThreeCard(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 8, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 5, Suit: 1},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 5, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 9, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 5, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 12, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
			},
			[]int {0, 3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.THREE_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 0, Suit: 1},
						card.Card{Number: 0, Suit: 2},
						card.Card{Number: 1, Suit: 3},
						card.Card{Number: 5, Suit: 1},
					},
				},
			},
			[]int {3},
		},
	}

	exec(t, threeCard, patterns)
}

func TestTwoPair(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
			},
			[]int {0, 3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 3, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 4, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 7, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 10, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.TWO_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 8, Suit: 0},
						card.Card{Number: 8, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 0, Suit: 1},
					},
				},
			},
			[]int {3},
		},
	}

	exec(t, twoPair, patterns)
}

func TestOnePair(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.ONE_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 3, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 0, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.ONE_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 5, Suit: 1},
						card.Card{Number: 0, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
			},
			[]int {3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.ONE_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 0, Suit: 1},
						card.Card{Number: 9, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.ONE_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 5, Suit: 0},
						card.Card{Number: 5, Suit: 1},
						card.Card{Number: 0, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.ONE_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 0, Suit: 1},
						card.Card{Number: 9, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 6, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.ONE_PAIR,
					UsedCards: evaluator.Cards{
						card.Card{Number: 0, Suit: 0},
						card.Card{Number: 0, Suit: 1},
						card.Card{Number: 9, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 4, Suit: 1},
					},
				},
			},
			[]int {0},
		},
	}

	exec(t, onePair, patterns)
}


func TestHighCard(t *testing.T) {
	patterns := []TestStructure{
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.HIGH_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 6, Suit: 3},
						card.Card{Number: 8, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.HIGH_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 6, Suit: 3},
						card.Card{Number: 8, Suit: 1},
					},
				},
			},
			[]int {0, 3},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.HIGH_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 2, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 6, Suit: 3},
						card.Card{Number: 8, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.HIGH_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 3, Suit: 1},
						card.Card{Number: 5, Suit: 2},
						card.Card{Number: 6, Suit: 3},
						card.Card{Number: 8, Suit: 1},
					},
				},
			},
			[]int {0},
		},
		{
			evaluator.Roles{
				evaluator.Role{
					PlayerId: 0,
					Role: evaluator.HIGH_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 1, Suit: 0},
						card.Card{Number: 2, Suit: 1},
						card.Card{Number: 3, Suit: 2},
						card.Card{Number: 5, Suit: 3},
						card.Card{Number: 0, Suit: 1},
					},
				},
				evaluator.Role{
					PlayerId: 3,
					Role: evaluator.HIGH_CARD,
					UsedCards: evaluator.Cards{
						card.Card{Number: 12, Suit: 0},
						card.Card{Number: 11, Suit: 1},
						card.Card{Number: 10, Suit: 2},
						card.Card{Number: 9, Suit: 3},
						card.Card{Number: 7, Suit: 1},
					},
				},
			},
			[]int {0},
		},
	}
	exec(t, highCard, patterns)
}
