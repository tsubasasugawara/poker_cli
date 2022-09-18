package structs

type Player struct {
	Id				int
	Uuid			string
	Hand			[2]Card
	Stack			int
	BettingAmount	int
}
