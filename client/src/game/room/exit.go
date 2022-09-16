package room

import (
	"net/http"
	"bytes"
	"encoding/json"

	"poker/env"
)

/*
 * ルームから退出
 * @{param} uid string : ユーザID
 * @{param} roomId string : ルームID
 * @{return} string : ステータスメッセージ
 */
func Exit(uid, roomId string) (string) {
	body, _ := json.Marshal(Room{UserId: uid, RoomId: roomId})
	resp, err := http.Post(env.ROOT + "/room/exit", "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var j Message
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return ""
	}

	return j.Status
}
