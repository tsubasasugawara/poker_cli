package play

import (
	"poker/game"
	"poker/game/dealer"
	"poker/game/player"
)

type Room struct {
	Dealer			*dealer.Dealer
	Players			[]*player.Player
	State 			int //ゲームの進行状況を格納(プリフロップ, フロップ, ターン, リバー)
	ActionHistory	game.ActionHistory
	Rate			int // レート (100/200 → 200を代入)
	Finish			bool
}
