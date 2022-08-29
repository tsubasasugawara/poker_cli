package evaluator

import (
	// "log"

	"poker/poker/playing_cards/card"
	"poker/poker/player"
)

type CardMap [card.SuitNum + 1][card.CardsNum + 1]int

// 役を評価し、勝者のプレイヤーIDを返す
func Evaluator(players []player.Player, board[5]card.Card) {
	var points Points

	for _, p := range players {
		cards := join(p.Hand, board)
		cardMap := MakeCardMap(cards)

		var role = HIGH_CARD 			// 役ができているかどうか0
		var usedCards []card.Card 	// 役を形成するカード

		if role == HIGH_CARD {
			role, usedCards = StraightFlush(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = FourCard(cardMap)
		}

		if role == HIGH_CARD {
			// フルハウス
		}

		if role == HIGH_CARD {
			// フラッシュ
		}

		if role == HIGH_CARD {
			// ストレート
		}

		if role == HIGH_CARD {
			// スリーカード
		}

		if role == HIGH_CARD {
			// ツーペア
		}

		if role == HIGH_CARD {
			// ワンペア
		}

		points = append(
			points,
			Point{
				PlayerId: p.Id,
				Role: role,
				UsedCards: usedCards,
			},
		)
	}
}

func join(hand [2]card.Card, board[5]card.Card) []card.Card {
	var cards []card.Card

	for _, v := range hand {
		cards = append(cards, v)
	}

	for _, v := range board {
		cards = append(cards, v)
	}

	return cards
}

// カードの大小を比べ、大きい方を返す
func cardsMax(a, b int) int {
	if a == 0 {
		a = a + 13
	}
	if b == 0 {
		b = b + 13
	}

	if a < b {
		return b
	} else {
		return a
	}
}

func cardsMin(a, b int) int {
	if a == 0 {
		a = a + 13
	}
	if b == 0 {
		b = b + 13
	}

	if a < b {
		return a
	} else {
		return b
	}
}

// 行の最後は同じスートのカード枚数,
// 列の最後は同じ数のカード枚数
func MakeCardMap(cards []card.Card) CardMap {
	cardMap := CardMap{}
	for _, c := range cards {
		cardMap[c.Suit][c.Number] = 1
		cardMap[c.Suit][card.CardsNum] += 1
		cardMap[card.SuitNum][c.Number] += 1
	}
	return cardMap
}

// カードマップの中から最も大きい値を見つける
func maxCard(cardMap CardMap) card.Card {
	for i := card.CardsNum; i > 0; i-- {
		for j := 0; j < card.SuitNum; j++ {
			if cardMap[i % 13][j] == 1 {
				return card.Card{Number: (i % 13), Suit: j}
			}
		}
	}
	return card.Card{Number: -1, Suit: -1}
}

// 1番目の戻り値は、役
// 2番目の戻り値は、役を形成するカード
func StraightFlush(cardMap CardMap) (int, []card.Card) {

	for i := 0; i < card.SuitNum; i++ {
		cards := []card.Card{} // 役を形成するカードを格納

		if cardMap[i][card.CardsNum] < 5 {
			continue
		}

		for num := 13; num >= 0; num-- {
			if cardMap[i][num % 13] == 1 {
				cards = append(cards, card.Card{Number: (num % 13), Suit: i})
			} else {
				cards = []card.Card{}
			}

			if len(cards) == 5 {
				break
			}
		}

		if len(cards) == 5 {
			return STRAIGHT_FLASH, cards
		}
	}

	return HIGH_CARD, []card.Card{}
}

func FourCard(cardMap CardMap) (int, []card.Card) {
	cards := []card.Card{}

	for num := 13; num > 0; num-- {
		if cardMap[card.SuitNum][num % 13] == 4 {
			cardMap[card.SuitNum][num % 13] = 0

			for i := 0; i < 4; i++ {
				cards = append(cards, card.Card{Number: (num % 13), Suit: i})
			}

			cards = append(cards, maxCard(cardMap))

			return FOUR_CARD, cards
		}
	}

	return HIGH_CARD, cards
}
