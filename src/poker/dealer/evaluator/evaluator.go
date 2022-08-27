package evaluator

import (
	"sort"
	"log"

	"poker/poker/playing_cards/card"
	"poker/poker/player"
)

// 役を評価し、勝者のプレイヤーIDを返す
func Evaluator(players []player.Player, board[5]card.Card) {
	var points Points

	for _, p := range players {
		cards := sortCards(p.Hand, board)

		var point int 	// 役の合計値
		var role int // 役ができているかどうか

		if role == HIGH_CARD {
			point, role = royalStraightFlush(cards)
		}

		if role == HIGH_CARD {
			// フォーカード
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
				Point: point,
				Role: role,
				HighCard: cardsMax(p.Hand[0].Number, p.Hand[1].Number),
			},
		)
	}

	log.Println(points)
}

func sortCards(hand [2]card.Card, board[5]card.Card) Cards {
	var cards Cards

	for _, v := range hand {
		cards = append(cards, v)
	}

	for _, v := range board {
		cards = append(cards, v)
	}

	sort.Sort(cards)

	return cards
}

// カードの大小を比べ、大きい方を返す
func cardsMax(a, b int) int {
	if a == 1 {
		a = a + 13
	}
	if b == 1 {
		b = b + 13
	}

	if a < b {
		return b
	} else {
		return a
	}
}

// ロイヤルストレートフラッシュ判定
// 役ができていたら2番目の引数がtrueになり、
// カードのナンバーの合計を返す
func royalStraightFlush(cards []card.Card) (int, int) {
	role := []card.Card{}

	for i := 10; i <= 14; i++ {
		for _, ele := range cards {
			if ((ele.Number % 13) == (i % 13)) {
				if len(role) <= 0 {
					role = append(role, ele)
				} else if role[0].Suit == ele.Suit {
					role = append(role, ele)
				}
				break
			}
		}
	}

	if len(role) == 5 {
		return (10 + 11 + 12 + 13 + 14), ROYAL_STRAIGHT_FLASH
	}
	return 0, HIGH_CARD
}

// フォーカード判定
// func fourCard(cards []card.Card) (int, int) {
// 	for _, v := range cards
// }
