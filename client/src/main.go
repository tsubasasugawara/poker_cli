package main

import (
    // "net/url"
	"encoding/json"
	"net/http"
	"bytes"
	"log"
	// "poker/game"
	"fmt"
	"bufio"
	"os"

	// "poker/terminal"
	"poker/env"
	"poker/terminal/utils"
	"poker/game/user"
	"poker/game/room"
)

var (
	USER_ID = ""
	ROOM_ID = ""
)

type Data struct {
	UserId string `json:UuserId"`
	RoomId string `json:RroomId"`
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	USER_ID = user.Pre(scanner)
	utils.Clear()

	running := true
	for running {
		scanner.Scan()
		switch scanner.Text() {
		case "p":
			roomId, _ := room.Participate(USER_ID, scanner)

			// roomから退出する
			body, err := json.Marshal(
				&Data{
						UserId: USER_ID,
						RoomId: roomId,
				},
			)
			if err != nil {
				log.Println("json error: ", err)
			}

			resp, err := http.Post(env.ROOT + "/room/exit", "application/json; charset=UTF-8", bytes.NewBuffer(body))
			if err != nil {
				log.Println("post error: ", err)
			}
			defer resp.Body.Close()

		case "c":
			ROOM_ID := room.Create(USER_ID, scanner)
			fmt.Printf("Your room id is \"%s\". \n", ROOM_ID)

		case "e":
			running = false
		}

		utils.Clear()
	}


}
