package judge

type Ids []int

func (i Ids) Len() int {
	return len(i)
}

func (i Ids) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Ids) Less(a, b int) bool {
	return i[a] < i[b]
}
