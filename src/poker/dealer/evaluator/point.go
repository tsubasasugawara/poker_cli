package evaluator

type Point struct {
	PlayerId int
	Role int
	Point int
	HighCard int
}

type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	if p[i].Role == p[j].Role {
		if p[i].Point == p[j].Point {
			return p[i].HighCard < p[j].HighCard
		}
		return p[i].Point < p[j].Point
	}

	return p[i].Role < p[j].Role
}
