package dealer

type ActionHistory struct {
	action int
	chip int
	player int
}

const (
	ERROR	int = 0
	FOLD 	int = 1
	CHECK 	int = 2
	CALL 	int = 3
	BET 	int = 4
	RAISE 	int = 5
	ALLIN	int = 6

	// 上乗せするたびにRERAISEを足し算する
	RERAISE int = 10
)
