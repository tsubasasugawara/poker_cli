package drawing

import (
	"fmt"

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

/*
 * ターミナルに描写する文字列を生成する
 * @{param} players []*player.Player
 * @{param} dealer.Board 	[5]card.Card
 * @{param} uid 	string 			 プレイヤーのUUID
 * @{param} winner 	[]int 			 勝敗がついたら勝者のID
 * @{result} 		string			 ターミナルに表示する文字列
 */
func Drawing(players []*player.Player, dealer dealer.Dealer, uid string, winner []int) (string) {
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

	return fmt.Sprintf(
		template,
		dealer.Pot, dealer.CurrentPlayer,
		p1.Id, p1.Stack, p1.BettingAmount, p1BTN,
		dealer.Board[0].Number + 1,
		dealer.Board[1].Number + 1,
		dealer.Board[2].Number + 1,
		dealer.Board[3].Number + 1,
		dealer.Board[4].Number + 1,
		suitIntToString(dealer.Board[0].Suit),
		suitIntToString(dealer.Board[1].Suit),
		suitIntToString(dealer.Board[2].Suit),
		suitIntToString(dealer.Board[3].Suit),
		suitIntToString(dealer.Board[4].Suit),
		p2.Id, p2.Stack, p2.BettingAmount, p2BTN,
		p2.Hand[0].Number + 1, suitIntToString(p2.Hand[0].Suit),
		p2.Hand[1].Number + 1, suitIntToString(p2.Hand[1].Suit),
	)
}
