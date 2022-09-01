package dealer

import (
	"poker/playing_cards/card"
)

type DealerInterface interface {
	/* 初期処理 */
	init()

	/*
	 * プレイヤーの参加処理
	 * @{result}]:プレイヤーID
	 * @{result}:満席の場合、エラーを返す
	 */
	AddPlayer() (int, error)

	/*
	 * プレイヤーの退出処理
	 * @{param}:プレイヤーID
	 */
	TakePlayer(int)

	/* カードをシャッフル */
	Shuffle()

	/*
	 * ポット計算
	 * @{param}:ベット金額
	 */
	CalcPot(int)

	/*
	 * 参加プレイヤー数を求める
	 * @{result}:参加プレイヤー数
	 */
	countPlayers() (int)

	/* 3人以上のとき、間に空席があるとボタンの位置がずれるバグ
	 * ボタンがどの位置かを求める
	 * @{param}:参加プレイヤー数
	 * @{result}:ボタンのインデックス
	 * {result}:1人以下の場合にエラーを変えす
	 */
	calcBtnPosition(int) (int, error)

	/* 3人以上のとき、間に空席があると空席にカードを配ってしまう
	 * ハンドを配る
	 * @{result}:プレイヤーの数の文だけ
	 * @{result}:1人以下の場合にエラーを返す
	*/
	Deal() ([][2]card.Card, error)

	/*
	 * デッキの先頭を取り出す
	 * @{result}:デッキの先頭のカード
	*/
	NextCard() (card.Card)
}
