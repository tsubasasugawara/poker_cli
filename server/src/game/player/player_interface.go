package player

type PlayerInterface interface {
	/*
	 * ベット処理
	 * @{param}:ベット金額
	 * @{resutl}:アクション
	 * @{result}:かけたチップ数
	*/
	Bet(int) (int, int)

	/*
	 * スタックを計算する
	 * @{param}:勝ち額
	*/
	Win(int)
}
