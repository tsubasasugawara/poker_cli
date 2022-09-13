package main

import (
    // "net/url"
	// "encoding/json"
	// "net/http"
	// "poker/game"
	"bufio"
	"os"
	"poker/terminal"
	"poker/terminal/utils"
	"poker/game/user"
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
			terminal.Run()
		case "exit":
			running = false
		}
		utils.Clear()
	}


}
