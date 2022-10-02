package play

import (
	"errors"
	"log"

	"poker/game"
	"poker/game/dealer/evaluator"
	"poker/controller/play/util"
	"poker/game/dealer/judge"
)

// 1つのボードが終わるごとに実行
func Init(h *Hub, roomId string, winner []int) {
	// 勝者にポットを渡す
	for _, v := range winner {
		h.rooms[roomId].Players[v].CalcStack(h.rooms[roomId].Dealer.Pot / len(winner))
	}

	h.rooms[roomId].Players[0].ResetBettingAmount()
	h.rooms[roomId].Players[1].ResetBettingAmount()

	// ディーラーの初期化
	h.rooms[roomId].Dealer.Init()

	// アクションの履歴をクリア
	h.rooms[roomId].ActionHistory = game.ActionHistory{}

	// ゲームの進行状況をプリフロップへ戻す
	h.rooms[roomId].State = PRE_FROP

	//ボタンを移動する
	h.rooms[roomId].Dealer.NextGame()
}

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
 * @{return} bool : カードを見せる(true), 隠す(false)
 * @{return} error
 */
func GameProgress(h *Hub, userAction Action) ([]int, bool, error) {
	// 人数がたりない
	if len(h.rooms[userAction.RoomId].Players) != 2 {
		return []int{}, false, errors.New("Not enough players.")
	}
	// もし現在アクションしていいプレイヤーでなければ何もしない
	if h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.CurrentPlayer].Uuid != userAction.UserId && userAction.ActionType != game.DEAL {
		return []int{}, false, errors.New("Illegal Player.")
	}

	var (
		err error
	)
	switch userAction.ActionType {
	case game.FOLD:
		h.rooms[userAction.RoomId].State = PRE_FROP

		// ポットに追加
		h.rooms[userAction.RoomId].Dealer.CalcPot(h.rooms[userAction.RoomId].Players[0].BettingAmount)
		h.rooms[userAction.RoomId].Dealer.CalcPot(h.rooms[userAction.RoomId].Players[1].BettingAmount)

		// プレイヤーのベットをリセット
		h.rooms[userAction.RoomId].Players[0].ResetBettingAmount()
		h.rooms[userAction.RoomId].Players[1].ResetBettingAmount()

		return []int{1 - util.GetPlayerIndex(h.rooms[userAction.RoomId].Players, userAction.UserId)}, false, nil
	default:
		_, err = h.Actions(userAction)
		if err != nil {
			log.Println(err)
			return []int{}, false, err
		}
	}

	ok := whetherNextState(h, userAction)

	// オールインの場合は無理やり進める
	// TODO : オールインの際はカードをゆっくりめくる
	if (h.rooms[userAction.RoomId].Players[0].Stack == 0 || h.rooms[userAction.RoomId].Players[1].Stack == 0) && ok {
		switch h.rooms[userAction.RoomId].State {
		case PRE_FROP:
			frop(h, userAction)
			turn(h, userAction)
			river(h, userAction)
		case FROP:
			turn(h, userAction)
			river(h, userAction)
		case TURN:
			river(h, userAction)
		}

		h.rooms[userAction.RoomId].State = RIVER + 1
	}

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
		case RIVER:
			h.rooms[userAction.RoomId].State += 1
		}

		// アクション履歴の初期化
		h.rooms[userAction.RoomId].ActionHistory = game.ActionHistory{}

		// 最初にアクションするプレイヤーの設定
		h.rooms[userAction.RoomId].Dealer.FirstPlayer()

	}

	// 勝敗をジャッジする
	winner := []int{}
	show := false
	if h.rooms[userAction.RoomId].State == RIVER + 1 {
		roles := evaluator.Evaluator(h.rooms[userAction.RoomId].Players, h.rooms[userAction.RoomId].Dealer.Board)
		winner = judge.Judge(roles)

		// 勝敗が決まったらカードをショー
		if len(winner) > 0 {
			show = true
		}
	}

	return winner, show, nil
}
