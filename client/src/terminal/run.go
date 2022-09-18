package terminal

import (
	"log"
	"sync"

	"github.com/nsf/termbox-go"
	"github.com/gorilla/websocket"
)

func Run(uid, roomId string, conn *websocket.Conn, swg *sync.WaitGroup) {
	defer swg.Done()

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	GetChar(uid, roomId, conn)
}
