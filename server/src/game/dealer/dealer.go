package dealer

import (
	"math/rand"
	"time"
	"errors"

	"poker/game/playing_cards/card"
)

type Dealer struct {
	CurrentPlayer int // アクション待ちのプレイやー
	BigBlindPosition int
	Cards []card.Card // デッキ
	Board [5]card.Card // ボード上のカード
	Pot int
}

// コンストラクタ
func NewDealer() *Dealer {
	dealer := Dealer{}
	dealer.init()
	return &dealer
}

func (dealer *Dealer) init() {
	dealer.Board = [5]card.Card{
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
	}
	dealer.Pot = 0
	dealer.BigBlindPosition = 0

	dealer.Cards = []card.Card{}
	// 順番に並んだトランプを生成
	for i := 0; i < card.SuitNum; i++ {
		for j := 0; j < card.CardsNum; j++ {
			dealer.Cards = append(dealer.Cards, card.Card{Number: j, Suit: i})
		}
	}
}

func (dealer *Dealer) Shuffle() {
	dealer.init()

	rand.Seed(time.Now().UnixNano())

	for i, _ := range dealer.Cards {
		index := rand.Intn(card.SuitNum * card.CardsNum)
		dealer.Cards[index],dealer.Cards[i]  = dealer.Cards[i],dealer.Cards[index]
	}
}

// ゲームが次の段階へ進んだタイミングでポット計算を行う
func (dealer *Dealer) CalcPot(chip int) {
	dealer.Pot = dealer.Pot + chip
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

// 次のゲームへ進む
func (dealer *Dealer) NextGame(playersCnt int) {
	dealer.BigBlindPosition = (dealer.BigBlindPosition + 1) % playersCnt
}

// 最初にアクションをするプレイやーを指定する
func (dealer *Dealer) FirstPlayer(playersCnt int){
	dealer.CurrentPlayer = (dealer.BigBlindPosition + 1) % playersCnt
}

// 次にアクションをするプレイヤーを指定する
func (dealer *Dealer) NextPlayer(playersCnt int) {
	dealer.CurrentPlayer = (dealer.CurrentPlayer + 1) % playersCnt
}

func (dealer *Dealer) Deal(playersCnt int) ([][2]card.Card, error){
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
		res[position][int(i / playersCnt)] = dealer.NextCard()
	}

	// 最初にアクションをするプレイヤーを決定
	dealer.FirstPlayer(playersCnt)

	return res, nil
}

func (dealer *Dealer) NextCard() (card.Card) {
	res := dealer.Cards[0]
	dealer.Cards = dealer.Cards[1:]
	return res
}
