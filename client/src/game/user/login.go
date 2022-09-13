package user

import (
	"net/http"
	"bytes"
	"encoding/json"

	"poker/env"
)

/*
 * ログイン処理
 * @{param} name string : ユーザ名
 * @{param} password string
 * @{return} string : ユーザID
 */
func login(name, password string) (string) {
	body, _ := json.Marshal(User{Name: name, Password: password})
	resp, err := http.Post(env.ROOT + "/user/login", "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var j Message
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return ""
	}

	return j.Id
}
