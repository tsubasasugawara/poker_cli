package user

import (
	"net/http"
	"bytes"
	"encoding/json"
	"errors"

	"poker/env"
)
/*
 * ユーザ登録処理
 * @{param} name string : ユーザ名
 * @{param} password string
 * @{return} error
 */
func signup(name, password string) (error) {
	body, _ := json.Marshal(User{Name: name, Password: password})
	resp, err := http.Post(env.ROOT + "/user/regist", "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var j Message
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return err
	}

	if j.Status == SUCCESS {
		return nil
	} else {
		return errors.New(FAILURE)
	}
}
