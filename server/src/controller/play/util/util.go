package util

import (
	"poker/game/player"
)

func GetPlayerIndex(players []*player.Player, uid string) (int) {
	playerIndex := 0
	for i, p := range players {
		if uid == p.Uuid {
			playerIndex = i
		}
	}
	return playerIndex
}
