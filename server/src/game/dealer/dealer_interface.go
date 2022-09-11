package dealer

import (
	"poker/game/playing_cards/card"
)

type DealerInterface interface {
	/* 初期処理 */
	init()

	/* カードをシャッフル */
	Shuffle()

	/*
	 * ポット計算
	 * @{param}:ベット金額
	 */
	CalcPot(int)

	/* 3人以上のとき、間に空席があるとボタンの位置がずれるバグ
	 * ボタンがどの位置かを求める
	 * @{param}:参加プレイヤー数
	 * @{result}:ボタンのインデックス
	 * {result}:1人以下の場合にエラーを変えす
	 */
	calcBtnPosition(int) (int, error)

	/*
	 * 次のゲームへ移行する処理(BBをずらす)
	 */
	NextGame()

	/*
	 * 最初にアクションをするプレイヤー
	 */
	FirstPlayer()

	/*
	 * 次にアクションをするプレイヤーを設定
	 */
	NextPlayer()

	/* 3人以上のとき、間に空席があると空席にカードを配ってしまう
	 * ハンドを配る
	 * @{param}: プレイヤー数
	 * @{result}:プレイヤーの数の文だけ
	 * @{result}:1人以下の場合にエラーを返す
	*/
	Deal(plaerysCnt int) ([][2]card.Card, error)

	/*
	 * デッキの先頭を取り出す
	 * @{result}:デッキの先頭のカード
	*/
	NextCard() (card.Card)
}
