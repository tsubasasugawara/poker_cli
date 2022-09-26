package drawing

import (
	"fmt"
	"strconv"

	"poker/controller/play/util"
	"poker/game/player"
	"poker/game/dealer"
	"poker/game/playing_cards/suits"
)

func suitIntToString(suit int) string {
	var res string
	switch suit {
	case suits.HEART:
		res = "♥"
	case suits.DIAMOND:
		res = "◆"
	case suits.SPADE:
		res = "♤"
	case suits.CLUB:
		res = "♧"
	default:
		res = "?"
	}
	return res
}

func cardsNumToString(num int) string {
	num += 1

	var res string
	switch num {
	case 11:
		res = "J"
	case 12:
		res = "Q"
	case 13:
		res = "K"
	case 1:
		res = "A"
	case 0:
		res = "?"
	default:
		res = strconv.Itoa(num)
	}

	return res
}

/*
 * ターミナルに描写する文字列を生成する
 * @{param} players []*player.Player
 * @{param} dealer.Board 	[5]card.Card
 * @{param} uid 	string 			 プレイヤーのUUID
 * @{param} winner 	[]int 			 勝敗がついたら勝者のID
 * @{param} show 	bool 			 カードを見せるかどうか
 * @{result} 		string			 ターミナルに表示する文字列
 */
func Drawing(players []*player.Player, dealer dealer.Dealer, uid string, winner []int, show bool) (string) {
	if len(players) < 2 {
		return ""
	}

	idx := util.GetPlayerIndex(players, uid)

	p1 := players[1 - idx]
	p2 := players[idx]

	btn, err := dealer.CalcBtnPosition(len(players));
	if err != nil {
		return ""
	}

	p1BTN := ""
	p2BTN := ""
	if btn == idx {
		p2BTN = "BTN"
	} else {
		p1BTN = "BTN"
	}

	// 前から一枚目の番号,スート、二枚目の番号,スート
	opponentHand := [4]string{"?", "?", "?", "?"}

	if len(winner) > 0 && show {
		opponentHand[0] = cardsNumToString(p1.Hand[0].Number)
		opponentHand[1] = suitIntToString(p1.Hand[0].Suit)
		opponentHand[2] = cardsNumToString(p1.Hand[1].Number)
		opponentHand[3] = suitIntToString(p1.Hand[1].Suit)
	}

	return fmt.Sprintf(
		template,
		dealer.Pot, dealer.CurrentPlayer,
		p1.Id, p1.Stack, p1.BettingAmount, p1BTN,
		opponentHand[0], opponentHand[1], opponentHand[2], opponentHand[3],
		cardsNumToString(dealer.Board[0].Number),
		cardsNumToString(dealer.Board[1].Number),
		cardsNumToString(dealer.Board[2].Number),
		cardsNumToString(dealer.Board[3].Number),
		cardsNumToString(dealer.Board[4].Number),
		suitIntToString(dealer.Board[0].Suit),
		suitIntToString(dealer.Board[1].Suit),
		suitIntToString(dealer.Board[2].Suit),
		suitIntToString(dealer.Board[3].Suit),
		suitIntToString(dealer.Board[4].Suit),
		p2.Id, p2.Stack, p2.BettingAmount, p2BTN,
		cardsNumToString(p2.Hand[0].Number), suitIntToString(p2.Hand[0].Suit),
		cardsNumToString(p2.Hand[1].Number), suitIntToString(p2.Hand[1].Suit),
	)
}
