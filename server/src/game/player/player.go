package player

import (
	"poker/game"
	"poker/game/playing_cards/card"
)

type Player struct {
	Id				int				// ゲームで使うID
	Uuid			string			// ユーザのuuid
	Hand			[2]card.Card	// ハンド
	Stack			int				// 残りスタック数
	BettingAmount	int				// 現在のベット金額
	WinRecords		int				// 勝ち数
}

func NewPlayer(stack int, uuid string, id int) *Player {
	player := Player{Stack: stack, Uuid: uuid, Id: id}
	return &player
}

func (p *Player) Init(stack int) {
	p.Stack = stack
}

func (p *Player) Bet(chip int) (int, int) {
	if chip > p.Stack {
		return game.ERROR, 0
	} else if chip == p.Stack {
		p.Stack = 0
		return game.ALLIN, chip
	} else if chip > 0 {
		p.Stack = p.Stack - chip
		return game.BET, chip
	} else {
		return game.CHECK, chip
	}
}

func (p *Player) CalcStack(chip int) {
	p.Stack = p.Stack + chip
}

func (p *Player) ResetBettingAmount() {
	p.BettingAmount = 0
}

// ベット金額をBettingAmountに追加し、スタックから引く
func (p *Player) CalcBettingAmount(chip int) {
	p.BettingAmount += chip
	p.Stack -= chip
}
