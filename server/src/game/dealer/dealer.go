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
	dealer.Init()
	return &dealer
}

func (dealer *Dealer) Init() {
	dealer.Board = [5]card.Card{
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
		card.Card{Number: -1, Suit: -1},
	}
	dealer.Pot = 0
	dealer.NextGame()
	dealer.FirstPlayer()

	dealer.Cards = []card.Card{}
	// 順番に並んだトランプを生成
	for i := 0; i < card.SuitNum; i++ {
		for j := 0; j < card.CardsNum; j++ {
			dealer.Cards = append(dealer.Cards, card.Card{Number: j, Suit: i})
		}
	}
}

func (dealer *Dealer) Shuffle() {
	dealer.Init()

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

func (dealer *Dealer) CalcBtnPosition(playersCnt int) (int, error) {
	var btnPosition int
	if playersCnt <= 1 {
		return -1, errors.New("There are not enough players.")
	} else if playersCnt == 2 {
		btnPosition = 1 - dealer.BigBlindPosition
	} else {
		btnPosition = (dealer.BigBlindPosition - 2 + playersCnt) % 8
		return -1, errors.New("Too many players.")
	}

	return btnPosition, nil
}

// 次のゲームへ進む
func (dealer *Dealer) NextGame() {
	dealer.BigBlindPosition = 1 - dealer.BigBlindPosition
}

// 最初にアクションをするプレイやーを指定する
func (dealer *Dealer) FirstPlayer(){
	dealer.CurrentPlayer = 1 - dealer.BigBlindPosition
}

// 次にアクションをするプレイヤーを指定する
func (dealer *Dealer) NextPlayer() {
	dealer.CurrentPlayer = 1 - dealer.CurrentPlayer
}

func (dealer *Dealer) Deal(playersCnt int) ([][2]card.Card, error){
	var res [][2]card.Card
	for i := 0; i < playersCnt; i++ {
		res = append(res, [2]card.Card{})
	}

	// BTNの位置をBBから求める
	btnPosition, err := dealer.CalcBtnPosition(playersCnt)
	if err != nil {
		return res, err
	}

	// カードの割当
	for i := 0; i < playersCnt * 2; i++ {
		position := (btnPosition + 3 + i) % playersCnt
		res[position][int(i / playersCnt)] = dealer.NextCard()
	}

	// 最初にアクションをするプレイヤーを決定
	dealer.FirstPlayer()

	return res, nil
}

func (dealer *Dealer) NextCard() (card.Card) {
	res := dealer.Cards[0]
	dealer.Cards = dealer.Cards[1:]
	return res
}
