package structs

type Dealer struct {
	CurrentPlayer int // アクション待ちのプレイやー
	BigBlindPosition int
	Cards []Card // デッキ
	Board [5]Card // ボード上のカード
	Pot int
}
