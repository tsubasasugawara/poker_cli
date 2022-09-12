// 役ごとの勝者を判定する関数を使う際には、
// プレイヤーごとのカードをソートしてから
// 入力する必要がある。

package judge

import (
	"sort"
	"math"

	"poker/game/dealer/evaluator"
)

type Point struct {
	PlayerId int
	Point int
}

// 勝敗を判定し、勝者のプレイヤーIDを一覧で返す
func Judge(roles evaluator.Roles) []int {
	var winnersIds []int

	// 最も強い役を探しながら
	maxRole := evaluator.HIGH_CARD
	for _, v := range roles {
		if maxRole < v.Role {
			maxRole = v.Role
		}
	}

	// 最も強い役を持っているプレイヤーを探す
	var candidates evaluator.Roles
	for _, v := range roles {
		if maxRole == v.Role {
			candidates = append(candidates, v)
		}
	}

	if len(candidates) == 0{
		// ここでIDリストを返す
	}

	// 候補者の中から更に順位付けを行う
	switch maxRole {
		case evaluator.STRAIGHT_FLUSH:
			winnersIds = straightFlush(candidates)
		case evaluator.FOUR_CARD:
			winnersIds = fourCard(candidates)
		case evaluator.FULL_HOUSE:
			winnersIds = fullHouse(candidates)
		case evaluator.FLUSH:
			winnersIds = flush(candidates)
		case evaluator.STRAIGHT:
			winnersIds = straight(candidates)
		case evaluator.THREE_CARD:
			winnersIds = threeCard(candidates)
		case evaluator.TWO_PAIR:
			winnersIds = twoPair(candidates)
		case evaluator.ONE_PAIR:
			winnersIds = onePair(candidates)
		case evaluator.HIGH_CARD:
			winnersIds = highCard(candidates)
	}

	return winnersIds
}

// 勝者のリストを作成
func makeWinnerIdList(points []Point, maxPoint int) Ids {
	var winnerIds Ids
	for _, p := range points {
		if p.Point == maxPoint {
			winnerIds = append(winnerIds, p.PlayerId)
		}
	}
	return winnerIds
}

func aceToThirteen(num int) int {
	if num == 0 {
		num = 13
	}
	return num
}

// ストレート・フラッシュ同士の勝敗を判定し、勝者のIDを返す
func straightFlush(candidates evaluator.Roles) Ids{
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		point := 0

		// ロイヤルストレートフラッシュのときに、0を13にすることで強さを表現
		if c.UsedCards[4].Number == 0 {
			if c.UsedCards[3].Number == 12 {
				point = 13
			} else {
				point = c.UsedCards[3].Number
			}
		} else {
			point = c.UsedCards[4].Number
		}

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func fourCard(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		point := 0

		fourCardNumber := -1 //フォーカードを構成するカードの数
		highCardNumber := -1 // 残りの一枚の数

		// 最初の四枚がフォーカードのとき(ソートされている前提)
		if c.UsedCards[0].Number == c.UsedCards[1].Number && c.UsedCards[1].Number == c.UsedCards[2].Number && c.UsedCards[2].Number == c.UsedCards[3].Number {
			fourCardNumber = aceToThirteen(c.UsedCards[0].Number)
			highCardNumber = aceToThirteen(c.UsedCards[4].Number)
		} else {
			fourCardNumber = aceToThirteen(c.UsedCards[4].Number)
			highCardNumber = aceToThirteen(c.UsedCards[0].Number)
		}

		// 重み付けをすることでポイントを大小を求める
		point = fourCardNumber * 100 + highCardNumber
		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func fullHouse(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		threeCardNumber := -1 // スリーカードを構成するカードの番号
		pairNumber := -1		// ペアを構成するカードの番号

		// 最初の三枚がスリーカードのとき(ソートされている前提)
		if c.UsedCards[0].Number == c.UsedCards[1].Number && c.UsedCards[1].Number == c.UsedCards[2].Number {
			threeCardNumber = aceToThirteen(c.UsedCards[0].Number)
			pairNumber = aceToThirteen(c.UsedCards[3].Number)
		} else {
			threeCardNumber = aceToThirteen(c.UsedCards[2].Number)
			pairNumber = aceToThirteen(c.UsedCards[0].Number)
		}

		// 重み付けをすることでポイントの大小を求める
		point := threeCardNumber * 100 + pairNumber
		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func flush(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		point := 0

		for i := len(c.UsedCards) - 1; i >= 0; i-- {
			// 一番大きいカードに一番大きな重みをかける(ソートされている前提)
			point = point * 100 + c.UsedCards[i].Number
		}

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func straight(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		point := 0

		if c.UsedCards[4].Number == 0 {
			if c.UsedCards[3].Number == 12 {
				point = 13
			} else {
				point = c.UsedCards[3].Number
			}
		} else {
			point = c.UsedCards[4].Number
		}

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func threeCard(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)
		point := 0

		// ○○○□△ or □○○○△ or □△○○○
		if c.UsedCards[0].Number == c.UsedCards[1].Number && c.UsedCards[1].Number == c.UsedCards[2].Number {
			point = aceToThirteen(c.UsedCards[2].Number) * int(math.Pow(100, 2)) + aceToThirteen(c.UsedCards[4].Number) * 100 + aceToThirteen(c.UsedCards[3].Number)
		} else if c.UsedCards[1].Number == c.UsedCards[2].Number && c.UsedCards[2].Number == c.UsedCards[3].Number {
			point = aceToThirteen(c.UsedCards[2].Number) * int(math.Pow(100, 2)) + aceToThirteen(c.UsedCards[4].Number) * 100 + aceToThirteen(c.UsedCards[0].Number)
		} else {
			point = aceToThirteen(c.UsedCards[2].Number) * int(math.Pow(100, 2)) + aceToThirteen(c.UsedCards[1].Number) * 100 + aceToThirteen(c.UsedCards[0].Number)
		}

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func twoPair(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		// ○□□△△ or □□○△△ or □□△△○
		point := aceToThirteen(c.UsedCards[4].Number) * int(math.Pow(100, 2)) + aceToThirteen(c.UsedCards[2].Number) * 100 + aceToThirteen(c.UsedCards[0].Number)

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func onePair(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		// ペアを見つけだす
		i := 1
		for ; i < len(c.UsedCards); i++ {
			if c.UsedCards[i].Number == c.UsedCards[i - 1].Number {
				break
			}
		}

		point := aceToThirteen(c.UsedCards[i].Number)

		// 100倍することでポイントで勝敗を計算できるようにする
		for j := len(c.UsedCards) - 1; j >= 0; j-- {
			if j == i || j == i - 1 {
				continue
			}
			point = point * 100 +  aceToThirteen(c.UsedCards[j].Number)
		}

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}

func highCard(candidates evaluator.Roles) Ids {
	var points []Point
	var maxPoint int
	for _, c := range candidates {
		sort.Sort(c.UsedCards)

		point := 0

		// 100倍することでポイントで勝敗を計算できるようにする
		for i := len(c.UsedCards) - 1; i >= 0; i-- {
			point = point * 100 + aceToThirteen(c.UsedCards[i].Number)
		}

		if maxPoint < point {
			maxPoint = point
		}
		points = append(points, Point{PlayerId: c.PlayerId, Point: point})
	}

	return makeWinnerIdList(points, maxPoint)
}
