package play

import (
	"log"
	"time"

	"poker/game"
	"poker/game/player"
	"poker/game/dealer"
	"poker/game/drawing"
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
	Dealer 	dealer.Dealer 	`json:"dealer"`
	Players []player.Player	`json:"players"`
	Winner 	[]int			`json:"winner"`
	Msg		string			`json:"msg"`
	ErrMsg	error			`json:"errMsg"`
}

func NewHub() *Hub {
	return &Hub{
		broadcast:	make(chan Action, 10),
		register:	make(chan *Client, 10),
		unregister:	make(chan *Client, 10),
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

// Clientの接続情報を追加
func (h *Hub) addClient(roomId string, client *Client) {
	if _, exist := h.clients[roomId]; !exist {
    	h.clients[roomId] = make(map[*Client]bool)
	}
	h.clients[roomId][client] = true
}

// プレイヤーをルームに追加
func (h *Hub) addPlayer(roomId string, userId string, id int) {
	h.rooms[roomId].Players = append(h.rooms[roomId].Players, player.NewPlayer(5000, userId, id))
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.makeRoom(client.Info.RoomId)

			// ヘッズアップのみに制限
			if len(h.rooms[client.Info.RoomId].Players) >= 2 {
				continue
			}

			h.addClient(client.Info.RoomId, client)

			h.addPlayer(client.Info.RoomId, client.Info.UserId, len(h.rooms[client.Info.RoomId].Players))

			// 2人になったらゲームを開始する
			if len(h.rooms[client.Info.RoomId].Players) == 2 {
				h.rooms[client.Info.RoomId].Dealer = dealer.NewDealer()
				h.broadcast <- Action{UserId: client.Info.UserId, RoomId: client.Info.RoomId, ActionType: game.DEAL, Data: "Deal the cards."}
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client.Info.RoomId][client]; ok {
				delete(h.clients[client.Info.RoomId], client)
				//ルームの人数が0人だったらルーム削除
				if len(h.clients[client.Info.RoomId]) == 0 {
					delete(h.clients, client.Info.RoomId)
				}
			}

			client.conn.Close()

			// ルームからプレイヤーを削除
			for i, p := range h.rooms[client.Info.RoomId].Players {
				if p.Uuid == client.Info.UserId {
					h.rooms[client.Info.RoomId].Players = append(h.rooms[client.Info.RoomId].Players[:i], h.rooms[client.Info.RoomId].Players[i+1:]...)
				}
			}

		case userAction := <-h.broadcast:
			if h.rooms[userAction.RoomId].Finish {
				continue
			}

			winner, show, err := GameProgress(h, userAction)
			if err != nil {
				log.Println(err)
			}

			// 同じルームのユーザにデータを送信
			for client := range h.clients[userAction.RoomId] {
				var players []player.Player
				for _, player := range h.rooms[userAction.RoomId].Players {
					players = append(players, *player)
				}
				msg := drawing.Drawing(
					h.rooms[userAction.RoomId].Players,
					*h.rooms[userAction.RoomId].Dealer,
					client.Info.UserId,
					winner,
					show,
				)

				// チャネルにデータを送る
				select {
				case client.send <- msg:
				default:
					log.Println("default")
				}
			}

			// 勝敗が決まったら、3秒後に新しいボードへと移行する
			// TODO : 片方のスタックが0になったときにゲームを終了するかやり直すか
			// TODO : 同じ処理があるためリファクタリング必要
			if len(winner) > 0 {
				// 処理の開始
				h.rooms[userAction.RoomId].Finish = true

				time.Sleep(3 * time.Second)

				Init(h, userAction.RoomId, winner)

				h.rooms[userAction.RoomId].Dealer.Init()

				// 処理の終了
				h.rooms[userAction.RoomId].Finish = false

				// ディール
				h.broadcast <- Action{UserId: userAction.UserId, RoomId: userAction.RoomId, ActionType: game.DEAL, Data: "Deal the cards."}

			}
		}
	}
}
