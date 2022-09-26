package play

import (
	"strconv"
	"errors"

	"poker/game"
	"poker/controller/play/util"
)

type Action struct {
	UserId		string	`json:"userId"`
	RoomId		string	`json:"roomId"`
	ActionType	int		`json:"actionType`
	Data		string	`json:"data"`
}

/*
 * カードを配る
 * @{param} h *Hub
 * @{param} userAction Action
 * @{resutl} bool : オールインしたかどうか
 * @{resutl} error
*/
func deal(h *Hub, userAction Action) (bool, error) {
	// ディールの合図が２つ以上来た場合
	if h.rooms[userAction.RoomId].State != PRE_FROP {
		return false, nil
	}

	// ディーラーを初期化
	h.rooms[userAction.RoomId].Dealer.Init()

	//カードシャッフル
	h.rooms[userAction.RoomId].Dealer.Shuffle()

	// ハンドを配る
	playersCnt := len(h.rooms[userAction.RoomId].Players)
	hands, err := h.rooms[userAction.RoomId].Dealer.Deal(playersCnt)
	if err != nil {
		return false, err
	}
	h.rooms[userAction.RoomId].Players[0].Hand = hands[0]
	h.rooms[userAction.RoomId].Players[1].Hand = hands[1]

	// プリフロップへ
	// h.rooms[userAction.RoomId].State = PRE_FROP
	rate := 200
	h.rooms[userAction.RoomId].Rate = rate

	var (
		bbBettingAmount int = rate
		sbBettingAmount int = rate / 2

		bbAllIn bool = false
		sbAllIn bool = false
	)

	// BBが1BB以上スタックを持っているかどうかを確認し、
	// 持っていなかったらオールインとする
	if stack := h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.BigBlindPosition].Stack; stack <= bbBettingAmount {
		bbBettingAmount = stack
		bbAllIn = true
	}

	// SBが0.5BB以上スタックを持っているかどうかを確認し、
	// 持っていなかったらオールインとする
	if stack := h.rooms[userAction.RoomId].Players[1 - h.rooms[userAction.RoomId].Dealer.BigBlindPosition].Stack; stack <= sbBettingAmount {
		sbBettingAmount = stack
		sbAllIn = true
	}

	h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.BigBlindPosition].CalcBettingAmount(bbBettingAmount)
	h.rooms[userAction.RoomId].Players[1 - h.rooms[userAction.RoomId].Dealer.BigBlindPosition].CalcBettingAmount(sbBettingAmount)

	// BBが強制オールインのとき
	if bbAllIn {
		// BET扱いにする
		h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
			h.rooms[userAction.RoomId].ActionHistory,
			game.History{
				Action: game.BET,
				Chip: bbBettingAmount,
				PlayerId: h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.BigBlindPosition].Uuid,
			},
		)

		// SB(0.5BB)よりもBBのオールイン額が小さい場合はSBはコールとする
		h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
			h.rooms[userAction.RoomId].ActionHistory,
			game.History{
				Action: game.CALL,
				Chip: bbBettingAmount,
				PlayerId: h.rooms[userAction.RoomId].Players[1 - h.rooms[userAction.RoomId].Dealer.BigBlindPosition].Uuid,
			},
		)

		// チップを返す
		h.rooms[userAction.RoomId].Players[1 - h.rooms[userAction.RoomId].Dealer.BigBlindPosition].CalcBettingAmount(bbBettingAmount - sbBettingAmount)
	}

	// SBが強制オールインのとき
	if sbAllIn {
		// BET扱いにする
		h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
			h.rooms[userAction.RoomId].ActionHistory,
			game.History{
				Action: game.BET,
				Chip: sbBettingAmount,
				PlayerId: h.rooms[userAction.RoomId].Players[1 - h.rooms[userAction.RoomId].Dealer.BigBlindPosition].Uuid,
			},
		)

		// SB(0.5BB)よりもBBのオールイン額が小さい場合はSBはコールとする
		h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
			h.rooms[userAction.RoomId].ActionHistory,
			game.History{
				Action: game.CALL,
				Chip: bbBettingAmount,
				PlayerId: h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.BigBlindPosition].Uuid,
			},
		)

		// チップを返す
		h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.BigBlindPosition].CalcBettingAmount(sbBettingAmount - bbBettingAmount)
	}

	return (sbAllIn || bbAllIn), nil
}

/*
 * ベット処理
 * @{param} h *Hub
 * @{param} userAction Action
 * @{resutl} bool : オールインしたかどうか
 * @{resutl} error
*/
func bet(h *Hub, userAction Action) (bool, error) {
	chip, err := strconv.Atoi(userAction.Data)
	if err != nil {
		return false, errors.New("Illegal bet.")
	}

	// チップの指定した量がプレイヤーの持っている量よりも多い場合はオールインとする
	allIn := false
	playerIndex := util.GetPlayerIndex(h.rooms[userAction.RoomId].Players, userAction.UserId)
	if stack := h.rooms[userAction.RoomId].Players[playerIndex].Stack; stack < chip {
		chip = stack
		allIn = true
	}

	// チップが相手の掛け金の二倍ならベット、
	// 等しいならコール(ショートスタックのオールインも含む)、
	// それ以下ならエラー
	playerChip := h.rooms[userAction.RoomId].Players[playerIndex].BettingAmount + chip
	// もし相手のベット金額+スタックが、ベットしたプレイヤーよりも小さい場合には返金
	if playerChip >= h.rooms[userAction.RoomId].Players[1 - playerIndex].BettingAmount * 2 {
		h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
			h.rooms[userAction.RoomId].ActionHistory,
			game.History{Action: game.BET, Chip: playerChip, PlayerId: userAction.UserId},
		)
	} else if playerChip == h.rooms[userAction.RoomId].Players[1 - playerIndex].BettingAmount || allIn {
		return call(h, userAction)
	} else {
		return false, errors.New("Illegal bet.")
	}

	h.rooms[userAction.RoomId].Players[playerIndex].CalcBettingAmount(chip)
	h.rooms[userAction.RoomId].Dealer.NextPlayer()

	return allIn, nil
}

/*
 * コール処理
 * @{param} h *Hub
 * @{param} userAction Action
 * @{resutl} bool : オールインしたかどうか
 * @{resutl} error
*/
func call(h *Hub, userAction Action) (bool, error) {
	allIn := false

	// プリフロップ以外で最初にチップをかける行為はベットとする
	if h.rooms[userAction.RoomId].State != PRE_FROP && len(h.rooms[userAction.RoomId].ActionHistory) == 0 {
		return false, errors.New("Illegal call.")
	}

	// アクションをしたプレイヤーのインデックス取得
	playerIndex := util.GetPlayerIndex(h.rooms[userAction.RoomId].Players, userAction.UserId)

	// 相手のベット金額と自分のベット金額の差額
	chip := h.rooms[userAction.RoomId].Players[1 - playerIndex].BettingAmount - h.rooms[userAction.RoomId].Players[playerIndex].BettingAmount

	// スタックがコールに必要なチップ数残っていなければ、オールインとする
	if h.rooms[userAction.RoomId].Players[playerIndex].Stack < chip {
		chip = h.rooms[userAction.RoomId].Players[playerIndex].Stack
		allIn = true
	}

	h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
		h.rooms[userAction.RoomId].ActionHistory,
		game.History{Action: game.CALL, Chip: chip, PlayerId: userAction.UserId},
	)
	h.rooms[userAction.RoomId].Players[playerIndex].CalcBettingAmount(chip)
	h.rooms[userAction.RoomId].Dealer.NextPlayer()

	return allIn, nil
}

/*
 * チェック処理
 * @{param} h *Hub
 * @{param} userAction Action
 * @{resutl} error
*/
func check(h *Hub, userAction Action) (error) {
	// アクションをしたプレイヤーのインデックス取得
	playerIndex := util.GetPlayerIndex(h.rooms[userAction.RoomId].Players, userAction.UserId)

	// もし相手のベット金額と同じでない場合はチェックできない
	if h.rooms[userAction.RoomId].Players[playerIndex].BettingAmount != h.rooms[userAction.RoomId].Players[1 - playerIndex].BettingAmount {
		return errors.New("Illegal check.")
	}

	h.rooms[userAction.RoomId].ActionHistory = game.AppendActionHistory(
		h.rooms[userAction.RoomId].ActionHistory,
		game.History{Action: game.CHECK, Chip: 0, PlayerId: userAction.UserId},
	)
	h.rooms[userAction.RoomId].Dealer.NextPlayer()

	return nil
}

func (h *Hub) Actions(userAction Action) (bool, error) {
	var (
		err		error
		allIn	bool = false
	)

	switch userAction.ActionType {
	case game.DEAL:
		allIn, err = deal(h, userAction)
	case game.BET:
		allIn, err = bet(h, userAction)
	case game.CALL:
		allIn, err = call(h, userAction)
	case game.CHECK:
		err = check(h, userAction)
	}
	if err != nil {
		return false, err
	}

	return allIn, nil
}
