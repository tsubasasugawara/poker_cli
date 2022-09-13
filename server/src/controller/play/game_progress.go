package play

import (
	"errors"
	"log"

	"poker/game"
	"poker/game/dealer/evaluator"
	"poker/controller/play/util"
	"poker/game/dealer/judge"
)

/*
 * 次のステップへ進むかどうかを判定
 * @{param} h *Hub
 * @{param} userAction Action
 * @{result} bool : 進める場合はtrue
 */
func whetherNextState(h *Hub, userAction Action) (bool) {
	// ベット金額が揃っていなければ次へは進めない
	if h.rooms[userAction.RoomId].Players[0].BettingAmount != h.rooms[userAction.RoomId].Players[1].BettingAmount {
		return false
	}

	length := len(h.rooms[userAction.RoomId].ActionHistory)

	// アクションを2人とも行っていなければいけない
	if length < 2 {
		return false
	}

	// コールにチェックしたら次へ進む
	if h.rooms[userAction.RoomId].ActionHistory[length - 2].Action == game.CALL &&
		h.rooms[userAction.RoomId].ActionHistory[length - 1].Action == game.CHECK {
			return true
	}

	// ベットにコールしたら進む
	if h.rooms[userAction.RoomId].ActionHistory[length - 2].Action == game.BET &&
		h.rooms[userAction.RoomId].ActionHistory[length - 1].Action == game.CALL {
			return true
	}

	// チェックチェックで進む
	if h.rooms[userAction.RoomId].ActionHistory[length - 2].Action == game.CHECK &&
		h.rooms[userAction.RoomId].ActionHistory[length - 1].Action == game.CHECK {
			return true
	}

	return false
}

/*
 * フロップ(カードを三枚ボードにおく)
 */
func frop(h *Hub, userAction Action) (error) {
	if h.rooms[userAction.RoomId].State != PRE_FROP {
		return errors.New("Illegal state.")
	}

	h.rooms[userAction.RoomId].Dealer.NextCard()
	h.rooms[userAction.RoomId].Dealer.Board[0] = h.rooms[userAction.RoomId].Dealer.NextCard()
	h.rooms[userAction.RoomId].Dealer.Board[1] = h.rooms[userAction.RoomId].Dealer.NextCard()
	h.rooms[userAction.RoomId].Dealer.Board[2] = h.rooms[userAction.RoomId].Dealer.NextCard()

	h.rooms[userAction.RoomId].State = FROP

	return nil
}

/*
 * ターン(四枚目のカードを引く)
 */
func turn(h *Hub, userAction Action) (error) {
	if h.rooms[userAction.RoomId].State != FROP {
		return errors.New("Illegal state.")
	}

	h.rooms[userAction.RoomId].Dealer.NextCard()
	h.rooms[userAction.RoomId].Dealer.Board[3] = h.rooms[userAction.RoomId].Dealer.NextCard()

	h.rooms[userAction.RoomId].State = TURN

	return nil
}

/*
 * リバー(５枚目のカードを引く)
 */
func river(h *Hub, userAction Action) (error) {
	if h.rooms[userAction.RoomId].State != TURN {
		return errors.New("Illegal state.")
	}

	h.rooms[userAction.RoomId].Dealer.NextCard()
	h.rooms[userAction.RoomId].Dealer.Board[4] = h.rooms[userAction.RoomId].Dealer.NextCard()

	h.rooms[userAction.RoomId].State = RIVER

	return nil
}


/*
 * ゲーム進行
 * @{param} h *Hub
 * @{param} userAction Action
 * @{return} []int : 勝敗が決まったら勝者のインデックス番号を返す。決まらなければ長さ0
 */
func GameProgress(h *Hub, userAction Action) ([]int, error) {
	// もし現在アクションしていいプレイヤーでなければ何もしない
	if h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.CurrentPlayer].Uuid != userAction.UserId {
		return []int{}, errors.New("Illegal Player.")
	}

	var (
		allIn bool
		err error
	)
	switch userAction.ActionType {
	case game.FOLD:
		h.rooms[userAction.RoomId].State = PRE_FROP
		return []int{util.GetPlayerIndex(h.rooms[userAction.RoomId].Players, userAction.UserId)}, nil
	default:
		allIn, err = h.Actions(userAction)
		if err != nil {
			log.Println(err)
			return []int{}, err
		}
	}

	ok := whetherNextState(h, userAction)
	if ok {
		// ポッドにチップを追加
		h.rooms[userAction.RoomId].Dealer.CalcPot(h.rooms[userAction.RoomId].Players[0].BettingAmount)
		h.rooms[userAction.RoomId].Dealer.CalcPot(h.rooms[userAction.RoomId].Players[1].BettingAmount)

		// プレイヤーのベットをリセット
		h.rooms[userAction.RoomId].Players[0].ResetBettingAmount()
		h.rooms[userAction.RoomId].Players[1].ResetBettingAmount()

		switch h.rooms[userAction.RoomId].State {
		case PRE_FROP:
			frop(h, userAction)
		case FROP:
			turn(h, userAction)
		case TURN:
			river(h,userAction)
		}

		h.rooms[userAction.RoomId].State += 1
	}

	// オールインの場合は無理やり進める
	if allIn && ok {
		switch h.rooms[userAction.RoomId].State {
		case FROP:
			turn(h, userAction)
			river(h, userAction)
		case TURN:
			river(h, userAction)
		}

		h.rooms[userAction.RoomId].State = RIVER + 1
	}

	// 勝敗をジャッジする
	winner := []int{}
	if h.rooms[userAction.RoomId].State == RIVER + 1 {
		roles := evaluator.Evaluator(h.rooms[userAction.RoomId].Players, h.rooms[userAction.RoomId].Dealer.Board)
		winner = judge.Judge(roles)
		h.rooms[userAction.RoomId].State = PRE_FROP
	}

	return winner, nil
}
