package evaluator

const (
	ROYAL_STRAIGHT_FLUSH	int = 9
	STRAIGHT_FLUSH			int = 8
	FOUR_CARD				int = 7
	FULL_HOUSE				int = 6
	FLUSH					int = 5
	STRAIGHT				int = 4
	THREE_CARD				int = 3
	TWO_PAIR				int = 2
	ONE_PAIR				int = 1
	HIGH_CARD				int = 0
)

type Role struct {
	PlayerId int
	Role int
	UsedCards Cards
}

type Roles []Role

func (r Roles) Len() int {
	return len(r)
}

func (r Roles) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Roles) Less(i, j int) bool {
	return r[i].Role < r[j].Role
}

