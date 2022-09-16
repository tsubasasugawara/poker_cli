package main

import (
    // "net/url"
	// "encoding/json"
	// "net/http"
	// "poker/game"
	"fmt"
	"bufio"
	"os"
	// "poker/terminal"
	"poker/terminal/utils"
	"poker/game/user"
	"poker/game/room"
)

var (
	USER_ID = ""
	ROOM_ID = ""
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	USER_ID = user.Pre(scanner)
	utils.Clear()

	running := true
	for running {
		scanner.Scan()
		switch scanner.Text() {
		case "play":
			room.Participate(USER_ID, scanner)
		case "create":
			ROOM_ID := room.Create(USER_ID, scanner)
			fmt.Printf("Your room id is \"%s\". \n", ROOM_ID)
		case "exit":
			running = false
		}
		fmt.Println()
	}


}
