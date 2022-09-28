package room

import (
	"net/http"
	"bytes"
	"encoding/json"
	"bufio"
	"fmt"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"poker/terminal"
	"poker/terminal/utils"
	"poker/env"

	"github.com/gorilla/websocket"

)

type Action struct {
	UserId		string	`json:"userId"`
	RoomId		string	`json:"roomId"`
	ActionType	int		`json:"actionType`
	Data		string	`json:"data"`
}

/*
 * ルームに参加する
 * @{param} uid string : ユーザID
 * @{param} scanner *bufio.Scanner
 * @{return} string : 接続したルームのID
 * @{return} error
 */
func Participate(uid string, scanner *bufio.Scanner) (string, error) {
	fmt.Println("Please enter the room id.")
	fmt.Print("ID : ")
	scanner.Scan()
	roomId := scanner.Text()

	fmt.Println("Please enter the room password.")
	fmt.Print("Password : ")
	scanner.Scan()
	password := scanner.Text()

	body, _ := json.Marshal(Room{UserId: uid, RoomId: roomId, Password: password})
	resp, err := http.Post(env.ROOT + "/room/participate", "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var j Message
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return "", err
	}

	if j.Status == "success" {
		Connect(uid, roomId)
	}

	return roomId, nil
}

var addr = flag.String("addr", strings.Replace(env.ROOT, "http://", "", -1), "http service address")

/*
 * WebSocket通信でゲームに接続
 * @{param} uid string : ユーザID
 * @{param} roomId string :
 * @{param} uid string :
 */
func Connect(uid, roomId string) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/play"}
	log.Printf("connecting to %s", u.String())

	header := http.Header{}
	header.Add("userId", uid)
	header.Add("roomId", roomId)

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}

			utils.Clear()
			fmt.Println(string(msg))

			if err != nil {
				log.Println("read:", err)
				return
			}
		}
	}()

	terminal.Run(uid, roomId, c)
}
