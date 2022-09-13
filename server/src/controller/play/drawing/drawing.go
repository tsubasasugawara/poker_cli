package drawing

import (
	"fmt"

	"poker/game/playing_cards/card"
	"poker/controller/play/util"
	"poker/game/player"
)

/*
 * ターミナルに描写する文字列を生成する
 * @{param} players []*player.Player
 * @{param} board 	[5]card.Card
 * @{param} uid 	string 			 プレイヤーのUUID
 * @{param} winner 	[]int 			 勝敗がついたら勝者のID
 * @{result} 		string			 ターミナルに表示する文字列
 */
func Drawing(players []*player.Player, board [5]card.Card, uid string, winner []int) (string) {
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
		p1.Id, p1.Stack, p1.BettingAmount, p1Msg,
		board[0].Number, board[0].Suit, board[1].Number, board[1].Suit, board[2].Number, board[2].Suit, board[3].Number, board[3].Suit, board[4].Number, board[4].Suit,
		p2.Id, p2.Stack, p2.BettingAmount, p2Msg,
	)
}
