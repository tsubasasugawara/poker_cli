package play

import (
	"poker/game"
	"poker/game/player"
)

func getPlayerIndex(players []*player.Player, uid string) (int) {
	playerIndex := 0
	for i, p := range players {
		if uid == p.Uuid {
			playerIndex = i
		}
	}
	return playerIndex
}

func GameProgress(h *Hub, userAction Action) {
	// もし現在アクションしていいプレイヤーでなければ何もしない
	if h.rooms[userAction.RoomId].Players[h.rooms[userAction.RoomId].Dealer.CurrentPlayer].Uuid != userAction.UserId {
		return
	}

	switch userAction.ActionType {
	case game.FOLD:

	default:
		h.Actions(userAction)
	}
	// for client := range h.clients[userAction.RoomId] {
	// 	select {
	// 	case client.send <- userAction:
	// 	default:
	// 		close(client.send)
	// 		delete(h.clients[userAction.RoomId], client)
	// 	}
	// }
}
