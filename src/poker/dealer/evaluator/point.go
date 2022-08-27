package evaluator

type Point struct {
	PlayerId int
	Point int
}

type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	return p[i].Point < p[j].Point
}
