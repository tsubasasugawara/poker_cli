package game

import (
	"log"

	"poker/dealer"
	"poker/player"
	"poker/dealer/judge"
	"poker/dealer/evaluator"
)

func Start() {
	d := dealer.NewDealer()

	p1 := player.NewPlayer(1000)
	p2 := player.NewPlayer(1000)

	p1id, err1 := d.AddPlayer()
	if err1 != nil {
		return
	}
	p1.Id = p1id

	p2id,err2 := d.AddPlayer()
	if err2 != nil {
		return
	}
	p2.Id = p2id

	for i := 0; i < 10; i++ {
		d.Shuffle()
		res,_ := d.Deal()

		p1.Hand = res[0]
		p2.Hand = res[1]

		flop(d)
		turn(d)
		river(d)

		roles := evaluator.Evaluator(
			[]player.Player{*p1,*p2},
			d.Board,
		)

		log.Println(judge.Judge(roles))
	}

}

func flop(d *dealer.Dealer) {
	// バーンカードを捨てる
	d.NextCard()

	for i := 0; i < 3; i++ {
		d.Board[i] = d.NextCard()
	}
}

func turn(d *dealer.Dealer) {
	// バーンカードを捨てる
	d.NextCard()

	d.Board[3] = d.NextCard()
}

func river(d *dealer.Dealer) {
	// バーンカードを捨てる
	d.NextCard()

	d.Board[4] = d.NextCard()
}
