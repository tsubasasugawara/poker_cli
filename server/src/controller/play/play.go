package play

import (
	"log"

    "github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Echo(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}
	defer ws.Close()

	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write: ", err)
			break
		}
	}
}
