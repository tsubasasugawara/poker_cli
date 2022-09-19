package play

import (
	"log"
	"net/http"
	"time"
	"encoding/json"

    "github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space	= []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UserInfo struct {
	UserId string `json:"userId"`
	RoomId string `json:"roomId"`
}

type Client struct {
	hub		*Hub
	conn	*websocket.Conn
	send	chan string
	EnterAt	time.Time
	Info	UserInfo
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var action Action
		err = json.Unmarshal(msg, &action)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		c.hub.broadcast <- action
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		data, ok := <- c.send
		if !ok {
			err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
			}
			break
		}

		msg := []byte(data)
		log.Println(string(msg))

		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
	}
}

func ServeWs(hub *Hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn}
	client.Info.UserId = c.Request.Header.Get("userId")
	client.Info.RoomId = c.Request.Header.Get("roomId")
	client.EnterAt = time.Now()
	client.send = make(chan string, 5)
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
