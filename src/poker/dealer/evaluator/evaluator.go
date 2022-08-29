package evaluator

import (
	"poker/poker/playing_cards/card"
	"poker/poker/player"
)

type CardMap [card.SuitNum + 1][card.CardsNum + 1]int

// 役を評価し、勝者のプレイヤーIDを返す
func Evaluator(players []player.Player, board[5]card.Card) {
	var points Points

	for _, p := range players {
		cards := join(p.Hand, board)
		cardMap := makeCardMap(cards)

		var role = HIGH_CARD 			// 役ができているかどうか0
		var usedCards Cards 			// 役を形成するカード

		if role == HIGH_CARD {
			role, usedCards = straightFlush(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = fourCard(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = fullHouse(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = flash(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = straight(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = threeCard(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = twoPair(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = onePair(cardMap)
		}

		if role == HIGH_CARD {
			role, usedCards = highCard(cardMap)
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

func join(hand [2]card.Card, board[5]card.Card) Cards {
	var cards Cards

	for _, v := range hand {
		cards = append(cards, v)
	}

	for _, v := range board {
		cards = append(cards, v)
	}

	return cards
}

// 行の最後は同じスートのカード枚数,
// 列の最後は同じ数のカード枚数
func makeCardMap(cards Cards) CardMap {
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
			if cardMap[j][i % card.CardsNum] == 1{
				return card.Card{Number: (i % card.CardsNum), Suit: j}
			}
		}
	}
	return card.Card{Number: -1, Suit: -1}
}

// 7枚の中のカードかどうか
func isCard(cardMap CardMap, number, suit int) bool {
	if cardMap[suit][number % card.CardsNum] > 0{
		return true
	} else {
		return false
	}
}

// 1番目の戻り値は、役
// 2番目の戻り値は、役を形成するカード
func straightFlush(cardMap CardMap) (int, Cards) {

	for i := 0; i < card.SuitNum; i++ {
		cards := Cards{} // 役を形成するカードを格納

		if cardMap[i][card.CardsNum] < 5 {
			continue
		}

		for num := card.CardsNum; num >= 0; num-- {
			if cardMap[i][num % card.CardsNum] == 1 {
				cards = append(cards, card.Card{Number: (num % card.CardsNum), Suit: i})
			} else {
				cards = Cards{}
			}

			if len(cards) == 5 {
				break
			}
		}

		if len(cards) == 5 {
			return STRAIGHT_FLASH, cards
		}
	}

	return HIGH_CARD, Cards{}
}

func fourCard(cardMap CardMap) (int, Cards) {
	cards := findRoleWithSameCard(cardMap, 4)
	if len(cards) == 0 {
		return HIGH_CARD, cards
	} else {
		return FOUR_CARD, cards
	}
}

func fullHouse(cardMap CardMap) (int, Cards) {
	var tc, op Cards
	role := HIGH_CARD

	role, tc = threeCard(cardMap)
	if role != THREE_CARD {
		return HIGH_CARD, Cards{}
	}

	role, op = onePair(cardMap)
	if role != ONE_PAIR {
		return HIGH_CARD, Cards{}
	}

	return FULL_HOUSE, append(tc[0:3], op[0:2]...)
}

func flash(cardMap CardMap) (int, Cards) {
	var cards Cards

	for i := 0; i < card.SuitNum; i++ {
		if cardMap[i][card.CardsNum] < 5 {
			continue
		}

		for num := card.CardsNum; num > 0; num-- {
			if cardMap[i][num % card.CardsNum] != 1 {
				continue
			}

			cards = append(cards, card.Card{Number: num % card.CardsNum, Suit: i})
		}

		return FLASH, cards
	}

	return HIGH_CARD, Cards{}
}

func straight(cardMap CardMap) (int, Cards) {
	var cards Cards

	cnt := 0
	num := card.CardsNum
	for ; num >= 0; num-- {
		if cardMap[card.SuitNum][num % card.CardsNum] > 0 {
			cnt += 1
		} else {
			cnt = 0
		}

		if cnt == 5 {
			break
		}
	}

	if cnt != 5 {
		return HIGH_CARD, Cards{}
	}

	for i := num; i < num + 5; i++ {
		for s := 0; s < card.SuitNum; s++ {
			if cardMap[s][i % card.CardsNum] == 1 {
				cards = append(cards, card.Card{Number: i % card.CardsNum, Suit: s})
				break
			}
		}
	}

	return STRAIGHT, cards
}

func threeCard(cardMap CardMap) (int, Cards) {
	cards := findRoleWithSameCard(cardMap, 3)
	if len(cards) == 0 {
		return HIGH_CARD, cards
	} else {
		return THREE_CARD, cards
	}
}

func twoPair(cardMap CardMap) (int, Cards) {
	var cards Cards

	cnt := 0
	for num := card.CardsNum; num > 0; num-- {
		if cardMap[card.SuitNum][num % card.CardsNum] == 2 {
			cnt += 1

			cardMap[card.SuitNum][num % card.CardsNum] = 0

			for i := 0; i < card.SuitNum; i++ {
				if isCard(cardMap, num, i) {
					cardMap[i][num % card.CardsNum] = 0
					cards = append(cards, card.Card{Number: (num % card.CardsNum), Suit: i})
				}
			}

			if cnt == 2 {
				cards = append(cards, maxCard(cardMap))
				return TWO_PAIR, cards
			}
		}
	}

	return HIGH_CARD, Cards{}
}

func onePair(cardMap CardMap) (int, Cards) {
	cards := findRoleWithSameCard(cardMap, 2)
	if len(cards) == 0 {
		return HIGH_CARD, cards
	} else {
		return ONE_PAIR, cards
	}
}

func highCard(cardMap CardMap) (int, Cards) {
	var cards Cards

	for num := card.CardsNum; num > 0; num-- {
		for i := 0; i < card.SuitNum; i++ {
			if cardMap[i][num % card.CardsNum] == 1 {
				cards = append(cards, card.Card{Number: num % 13, Suit: i})
			}

			if len(cards) == 5 {
				return HIGH_CARD, cards
			}
		}
	}

	return HIGH_CARD, Cards{}
}


func findRoleWithSameCard(cardMap CardMap, numberOfCardsNeededToMakeRole int) (Cards) {
	cards := Cards{}

	for num := card.CardsNum; num > 0; num-- {
		if cardMap[card.SuitNum][num % card.CardsNum] == numberOfCardsNeededToMakeRole {
			// 確認済みのため0枚とする
			cardMap[card.SuitNum][num % card.CardsNum] = 0

			for i := 0; i < card.SuitNum; i++ {
				if isCard(cardMap, num, i) {
					cardMap[i][num % card.CardsNum] = 0
					cards = append(cards, card.Card{Number: (num % card.CardsNum), Suit: i})
				}
			}

			for i := 0; i < 5 - numberOfCardsNeededToMakeRole; i++{
				tmp := maxCard(cardMap)
				cardMap[tmp.Suit][tmp.Number] = 0
				cards = append(cards, tmp)
			}

			return cards
		}
	}
	return Cards{}
}
