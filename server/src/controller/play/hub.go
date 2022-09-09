package play

import (
)

type Action struct {
	UserId		string	`json:"userId"`
	RoomId		string	`json:"roomId"`
	ActionType	string	`json:"actionType`
	Data		string	`json:"data"`
}

type Hub struct {
	// ルームIDでチャネルを分ける
	clients map[string]map[*Client]bool
	broadcast chan Action
	register chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:	make(chan Action),
		register:	make(chan *Client),
		unregister:	make(chan *Client),
		clients:	make(map[string]map[*Client]bool),
	}
}

func (h *Hub) addClient(roomId string, client *Client) {
	if _, exist := h.clients[roomId]; !exist {
    	h.clients[roomId] = make(map[*Client]bool)
	}
	h.clients[roomId][client] = true
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			for client := range h.clients[client.Info.RoomId] {
				msg := Action{UserId: client.Info.UserId, RoomId: client.Info.RoomId, ActionType: "JOIN", Data: "Some one join room."}
				client.send <- msg
			}

			h.addClient(client.Info.RoomId, client)

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
				msg := Action{UserId: client.Info.UserId, RoomId: client.Info.RoomId, ActionType: "JOIN", Data: "Some one leave room."}
				client.send <- msg
			}

			client.conn.Close()

		case userAction := <-h.broadcast:
			for client := range h.clients[userAction.RoomId] {
				select {
				case client.send <- userAction:
				default:
					close(client.send)
					delete(h.clients[userAction.RoomId], client)
				}
			}
		}
	}
}
