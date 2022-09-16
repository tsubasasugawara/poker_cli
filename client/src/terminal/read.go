package terminal

import (
	"syscall"
	"log"
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
			log.Println(betAmount)

			err := conn.WriteJSON(Action{
				UserId: userId,
				RoomId: roomId,
				ActionType: BET,
				Data: betAmount,
			})
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
			err := conn.WriteJSON(Action{
				UserId: userId,
				RoomId: roomId,
				ActionType: CALL,
				Data: "",
			})
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""
			log.Println("success");

		case A:
			err := conn.WriteJSON(Action{
				UserId: userId,
				RoomId: roomId,
				ActionType: CHECK,
				Data: "",
			})
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""

		case F:
			err := conn.WriteJSON(Action{
				UserId: userId,
				RoomId: roomId,
				ActionType: FOLD,
				Data: "",
			})
			if err != nil {
				log.Println("write close:", err)
				return
			}

			betAmount = ""

		default:
			if ZERO <= p[0] && NINE >= p[0] && len(betAmount) < 19 {
				betAmount = betAmount + string(p[0])
			}
			log.Println(p)
		}
	}
}
