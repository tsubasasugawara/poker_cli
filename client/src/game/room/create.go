package room

import (
	"net/http"
	"bytes"
	"encoding/json"
	"bufio"
	"fmt"

	"poker/env"
)

/*
 * ルームを作成する
 * @{param} uid string : 作成者のユーザID
 * @{param} scanner *bufio.Scanner
 * @{return} string : ルームID
 */
func Create(uid string, scanner *bufio.Scanner) (string) {
	fmt.Println("Please enter the password for your room.")
	fmt.Print("Password : ")
	scanner.Scan()
	password := scanner.Text()

	body, _ := json.Marshal(Room{UserId: uid, Password: password})
	resp, err := http.Post(env.ROOT + "/room/create", "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var j Room
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return ""
	}

	return j.RoomId
}
