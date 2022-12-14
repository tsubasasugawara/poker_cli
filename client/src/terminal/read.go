package terminal

import (
	"syscall"
	"log"
	"encoding/json"

	"github.com/gorilla/websocket"
)

const (
	ENTER 		byte = 13
	CtrlC 		byte = 3
	BackSpace	byte = 127

	C			byte = 99  // call
	A 			byte = 97  // check
	F			byte = 102 // fold

	ZERO		byte = 48
	NINE		byte = 57
)

// バッファの値を取得する
func readBuffer(bufCh chan []byte) {
	buf := make([]byte, 1024)
	for {
		if n, err := syscall.Read(syscall.Stdin, buf); err == nil {
			bufCh <- buf[:n]
		}
	}
}

// 文字列の入力を受け取る
// TODO : websockdet通信を終了したあと、少しの間入力を受け付けない期間がある
func GetChar(userId, roomId string, conn *websocket.Conn) {
	bufCh := make(chan []byte, 1)
	go readBuffer(bufCh)

	running := true
	var betAmount string

	for running {
		p := <-bufCh

		if len(p)!= 1 {
			continue
		}

		switch p[0] {
		case CtrlC:
			running = false
		case ENTER:
			msg, err := json.Marshal(
				Action{
					UserId: userId,
					RoomId: roomId,
					ActionType: BET,
					Data: betAmount,
				},
			)
			if err != nil {
				log.Println("json marshal error: ", err)
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""

		case BackSpace:
			if len(betAmount) > 0 {
				betAmount = betAmount[:len(betAmount) - 1]
			}

		case C:
			msg, err := json.Marshal(
				Action{
					UserId: userId,
					RoomId: roomId,
					ActionType: CALL,
					Data: "",
				},
			)
			if err != nil {
				log.Println("json marshal error: ", err)
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""

		case A:
			msg, err := json.Marshal(
				Action{
					UserId: userId,
					RoomId: roomId,
					ActionType: CHECK,
					Data: "",
				},
			)
			if err != nil {
				log.Println("json marshal error: ", err)
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""

		case F:
			msg, err := json.Marshal(
				Action{
					UserId: userId,
					RoomId: roomId,
					ActionType: FOLD,
					Data: "",
				},
			)
			if err != nil {
				log.Println("json marshal error: ", err)
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""

		default:
			if ZERO <= p[0] && NINE >= p[0] && len(betAmount) < 19 {
				betAmount = betAmount + string(p[0])
			}
		}
	}

	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure,""))
	if err != nil {
		log.Println("write close:", err)
	}
	return
}
