package dealer

import (
	"math/rand"
	"time"
	"errors"

	"poker/poker/playing_cards/card"
)

type Dealer struct {
	Players [2]bool // プレイヤーを管理
	BigBlindPosition int
	Cards []card.Card // デッキ
	Board [5]card.Card // ボード上のカード
	Pot int
	ActionHistory []ActionHistory
}

// コンストラクタ
func NewDealer() *Dealer {
	dealer := Dealer{}
	dealer.init()
	return &dealer
}

func (dealer *Dealer) init() {
	dealer.Board = [5]card.Card{
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
		card.Card{Number: 0, Suit: 0},
	}
	dealer.Pot = 0
	dealer.ActionHistory = []ActionHistory{}
	dealer.BigBlindPosition = 0

	dealer.Cards = []card.Card{}
	// 順番に並んだトランプを生成
	for i := 0; i < card.SuitNum; i++ {
		for j := 1; j <= card.CardsNum; j++ {
			dealer.Cards = append(dealer.Cards, card.Card{Number: j, Suit: i})
		}
	}
}

// プレイヤーの参加登録
func (dealer *Dealer) AddPlayer() (int, error) {
	for i, seated := range dealer.Players {
		if !seated {
			dealer.Players[i] = true
			return i, nil
		}
	}
	return -1, errors.New("席が空いていません。")
}

// プレイヤーの退出
func (dealer *Dealer) TakePlayer(playerId int) {
	dealer.Players[playerId] = false
}

func (dealer *Dealer) Shuffle() {
	dealer.init()

	rand.Seed(time.Now().Unix())

	for i, _ := range dealer.Cards {
		index := rand.Intn(card.SuitNum * card.CardsNum)
		dealer.Cards[index],dealer.Cards[i]  = dealer.Cards[i],dealer.Cards[index]
	}
}

func (dealer *Dealer) CalcPot(chip int) {
	dealer.Pot = dealer.Pot + chip
}

func (dealer *Dealer) countPlayers() int{
	cnt := 0
	for _, v := range dealer.Players {
		if v {
			cnt = cnt + 1
		}
	}
	return cnt
}

func (dealer *Dealer) calcBtnPosition(playersCnt int) (int, error) {
	var btnPosition int
	if playersCnt <= 1 {
		return -1, errors.New("プレイヤーが足りません。")
	} else if playersCnt == 2 {
		btnPosition = 1 - dealer.BigBlindPosition
	} else {
		btnPosition = (dealer.BigBlindPosition - 2 + playersCnt) % 8
	}

	return btnPosition, nil
}

// 人数が足りない場合は２つ目の引数にfalseを返す
func (dealer *Dealer) Deal() ([][2]card.Card, error){
	playersCnt := dealer.countPlayers()
	var res [][2]card.Card
	for i := 0; i < playersCnt; i++ {
		res = append(res, [2]card.Card{})
	}

	// BTNの位置をBBから求める
	btnPosition, err := dealer.calcBtnPosition(playersCnt)
	if err != nil {
		return res, err
	}

	// カードの割当
	for i := 0; i < playersCnt * 2; i++ {
		position := (btnPosition + 3 + i) % playersCnt
		res[position][int(i / playersCnt)] = dealer.Cards[0]
		dealer.Cards = dealer.Cards[1:]
	}

	return res, nil
}
