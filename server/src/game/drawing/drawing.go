package drawing

import (
	"fmt"

	// "poker/game/playing_cards/card"
	"poker/controller/play/util"
	"poker/game/player"
	"poker/game/dealer"
)

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

	p1Msg := ""
	p2Msg := ""

	if len(winner) == len(players) {
		p1Msg = "DRAW"
		p2Msg = "DRAW"
	}

	if len(winner) == 1 {
		if idx == winner[0] {
			p1Msg = "WIN"
			p2Msg = "LOSE"
		} else {
			p1Msg = "LOSE"
			p2Msg = "WIN"
		}
	}

	return fmt.Sprintf(
		template,
		dealer.Pot,
		p1.Id, p1.Stack, p1.BettingAmount, p1Msg,
		dealer.Board[0].Number, dealer.Board[0].Suit, dealer.Board[1].Number, dealer.Board[1].Suit, dealer.Board[2].Number, dealer.Board[2].Suit, dealer.Board[3].Number, dealer.Board[3].Suit, dealer.Board[4].Number, dealer.Board[4].Suit,
		p2.Id, p2.Stack, p2.BettingAmount, p2Msg,
	)
}
