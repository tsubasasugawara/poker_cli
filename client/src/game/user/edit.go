package user

import (
	"net/http"
	"bytes"
	"encoding/json"
	"errors"

	"poker/env"
)

type EditUser struct {
	OldName string `json:"oldName"`
	OldPassword string `json:"oldPassword"`
	NewName string `json:"newName"`
	NewPassword string `json:"newPassword"`
}

/*
 * ログイン処理
 * @{param} "oldName" : "古いユーザ名"
 * @{param} "oldPassword" : "古いユーザのパスワード"
 * @{param} "newName" : "新しいユーザ名"
 * @{param} "newPassword" : "新しいユーザのパスワード"
 * @{return} error
 */
func edit(oldName, oldPassword, newName, newPassword string) (error) {
	body, _ := json.Marshal(EditUser{OldName: oldName, OldPassword: oldPassword, NewName: newName, NewPassword: newPassword})
	resp, err := http.Post(env.ROOT + "/user/edit", "application/json; charset=UTF-8", bytes.NewBuffer(body))
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
