package user

import (
	"log"
	"net/http"
	"poker/model/users"

    "github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

/*
 * ユーザ登録
 * @{param} c *gin.Context
 * @{request json} "name" : "ユーザ名"
 * @{request json} "password" : "ユーザのパスワード"
 * @{response} json {"statusMessage" : "success" or "failure"}
*/
func Regist(c * gin.Context) {
	var user User
	c.BindJSON(&user)

	statusCode, err := users.Insert(user.Name, user.Password)
	if err != nil {
		log.Println(err)
		return
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}

/*
 * ログイン
 * @{param} c *gin.Context
 * @{request json} "name" : "ユーザ名"
 * @{request json} "password" : "ユーザのパスワード"
 * @{response} json {"statusMessage" : "success" or "failure"}
 * @{response} json {"id" : userID(36) or ""}
*/
func Login(c * gin.Context){
	var user User
	c.BindJSON(&user)

	id, err := users.Select(user.Name, user.Password)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = users.UpdateAccessDate(id)
	if err != nil {
		log.Println(err)
		return
	}
	var msg = "success"
	if id == "" {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
		"id": id,
	})
}

/*
 * ユーザ削除
 * @{param} c *gin.Context
 * @{request json} "name" : "ユーザ名"
 * @{request json} "password" : "ユーザのパスワード"
 * @{response} json {"statusMessage" : "success" or "failure"}
*/
func Delete(c * gin.Context){
	var user User
	c.BindJSON(&user)

	statusCode, err := users.Delete(user.Name, user.Password)
	if err != nil {
		log.Println(err)
		return
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}

/*
 * ユーザ情報の編集
 * @{param} c *gin.Context
 * @{request json} "oldName" : "古いユーザ名"
 * @{request json} "oldPassword" : "古いユーザのパスワード"
 * @{request json} "newName" : "新しいユーザ名"
 * @{request json} "newPassword" : "新しいユーザのパスワード"
 * @{response} json {"statusMessage" : "success" or "failure"}
*/
func Edit(c * gin.Context){
	var user struct {
		OldName string `json:"oldName"`
		OldPassword string `json:"oldPassword"`
		NewName string `json:"newName"`
		NewPassword string `json:"newPassword"`
	}
	c.BindJSON(&user)

	statusCode, err := users.Update(user.OldName, user.OldPassword, user.NewName, user.NewPassword)
	if err != nil {
		log.Println(err)
		return
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}
