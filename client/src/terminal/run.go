package terminal

import (
	"log"

	"github.com/nsf/termbox-go"
	"github.com/gorilla/websocket"
)

func Run(uid, roomId string, conn *websocket.Conn) {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	GetChar(uid, roomId, conn)
}
