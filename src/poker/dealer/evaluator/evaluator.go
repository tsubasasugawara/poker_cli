package evaluator

import (
	"sort"
	"log"

	"poker/poker/playing_cards/card"
	"poker/poker/player"
)



// 役を評価し、勝者のプレイヤーIDを返す
func Evaluator(players []player.Player, board[5]card.Card) {
	// points := [len(players)]int

	for _, p := range players {
		cards := sortCards(p.Hand, board)
		royalStraightFlush(cards)
	}
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

// ロイヤルストレートフラッシュ判定
// そうだったら2番目の引数がtrueになり、
// カードのナンバーの合計を返す
func royalStraightFlush(cards []card.Card) {
	log.Println(cards)
}
