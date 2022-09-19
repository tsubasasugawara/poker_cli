package game

type History struct {
	Action int
	Chip int
	PlayerId string
}

type ActionHistory []History

const (
	ERROR	int = 0

	FOLD 	int = 1
	CHECK 	int = 2
	CALL 	int = 3
	BET 	int = 4
	RAISE 	int = 5
	ALLIN	int = 6
	RERAISE int = 7

	DEAL	int = 8
	JOIN	int = 9
	LEAVE	int = 10

	WIN		int = 11
	LOSE	int = 12
	DRAW	int = 13

	HISTORY_MAX int = 10
)

func AppendActionHistory(ah ActionHistory, history History) ActionHistory {
	if len(ah) >= HISTORY_MAX {
		ah = append(ah[1:], history)
	} else {
		ah = append(ah, history)
	}
	return ah
}
