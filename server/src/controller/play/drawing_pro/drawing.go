package drawing

import (
	"fmt"

	"poker/game/playing_cards/card"
	"poker/controller/play"
	"poker/game/player"
)

/*
 * ターミナルに描写する文字列を生成する
 * @{param} players []*player.Player
 * @{param} board [5]card.Card
 * @{param} uid string : プレイヤーのUUID
 * @{result} string : ターミナルに表示する文字列
 */
func drawing(players []*player.Player, board [5]card.Card, uid string) (string) {
	idx := play.getPlayerIndex(players, uid)

	p1 := players[1 - idx]
	p2 := players[idx]

	return fmt.Sprintf(
		template,
		p1.Id, p1.Stack, p1.BettingAmount,
		board[0].Number, board[0].Suit, board[1].Number, board[1].Suit, board[2].Number, board[2].Suit, board[3].Number, board[3].Suit, board[4].Number, board[4].Suit,
		p2.Id, p2.Stack, p2.BettingAmount,
	)
}
