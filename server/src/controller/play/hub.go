package play

import (
	"log"
	"poker/game"
	"poker/game/player"
	"poker/game/dealer"
)

type Hub struct {
	// ルームIDでチャネルを分ける
	clients map[string]map[*Client]bool
	rooms map[string]*Room
	broadcast chan Action
	register chan *Client
	unregister chan *Client
}

type TransmissionData struct {
	dealer 	dealer.Dealer 	`json:"dealer"`
	players []player.Player	`json:"players"`
	winner 	[]int			`json:"winner"`
	errMsg	error			`json:"errMsg"`
}

func NewHub() *Hub {
	return &Hub{
		broadcast:	make(chan Action),
		register:	make(chan *Client),
		unregister:	make(chan *Client),
		clients:	make(map[string]map[*Client]bool),
		rooms:		make(map[string]*Room),
	}
}

func (h *Hub) makeRoom(roomId string) {
	if _, exist := h.rooms[roomId]; !exist {
		h.rooms[roomId] = &Room{
			Dealer: dealer.NewDealer(),
			Players: []*player.Player{},
			ActionHistory: game.ActionHistory{},
			Rate: 200,
		}
	}
}

func (h *Hub) addClient(roomId string, client *Client) {
	if _, exist := h.clients[roomId]; !exist {
    	h.clients[roomId] = make(map[*Client]bool)
	}
	h.clients[roomId][client] = true
}

func (h *Hub) addPlayer(roomId string, userId string) {
	h.rooms[roomId].Players = append(h.rooms[roomId].Players, player.NewPlayer(5000, userId))
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.makeRoom(client.Info.RoomId)
			// ヘッズアップのみに制限
			if len(h.rooms[client.Info.RoomId].Players) >= 2 {
				h.unregister <- client
			}

			for client := range h.clients[client.Info.RoomId] {
				msg := Action{UserId: client.Info.UserId, RoomId: client.Info.RoomId, Data: "Some one join room."}
				client.send <- msg
			}


			h.addClient(client.Info.RoomId, client)

			h.addPlayer(client.Info.RoomId, client.Info.UserId)

			// 2人になったらゲームを開始する
			if len(h.rooms[client.Info.RoomId].Players) == 2 {
				h.broadcast <- Action{UserId: client.Info.UserId, RoomId: client.Info.RoomId, ActionType: game.DEAL, Data: "Deal the cards."}
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client.Info.RoomId][client]; ok {
				delete(h.clients[client.Info.RoomId], client)
				//ルームの人数が0人だったらルーム削除
				if len(h.clients[client.Info.RoomId]) == 0 {
					delete(h.clients, client.Info.RoomId)
				}
				// close(client.send)
			}

			for client := range h.clients[client.Info.RoomId] {
				msg := Action{UserId: client.Info.UserId, RoomId: client.Info.RoomId, ActionType: game.LEAVE, Data: "Some one leave room."}
				client.send <- msg
			}

			client.conn.Close()

			// ルームからプレイヤーを削除
			for i, p := range h.rooms[client.Info.RoomId].Players {
				if p.Uuid == client.Info.UserId {
					h.rooms[client.Info.RoomId].Players = append(h.rooms[client.Info.RoomId].Players[:i], h.rooms[client.Info.RoomId].Players[i+1:]...)
				}
			}

		case userAction := <-h.broadcast:
			winner, err := GameProgress(h, userAction)
			if err != nil {
				log.Println("game progress error.")
				continue
			}

			for client := range h.clients[userAction.RoomId] {
				var data TransmissionData
				if err != nil {
					data = TransmissionData{errMsg: err}
				} else {
					data = TransmissionData{
						dealer: *h.rooms[userAction.RoomId].Dealer,
						players: []player.Player{*h.rooms[userAction.RoomId].Players[0], *h.rooms[userAction.RoomId].Players[1]},
						winner: winner,
					}
				}

				client.send <- data
			}
		}
	}
}
