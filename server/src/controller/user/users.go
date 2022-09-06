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

func Regist(c * gin.Context) {
	var user User
	c.BindJSON(&user)

	statusCode, err := users.Insert(user.Name, user.Password)
	if err != nil {
		log.Println(err)
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}

func Login(c * gin.Context){
	var user User
	c.BindJSON(&user)

	id, err := users.Select(user.Name, user.Password)
	if err != nil {
		log.Println(err)
	}

	_, err = users.UpdateAccessDate(id)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Delete(c * gin.Context){
	var user User
	c.BindJSON(&user)

	statusCode, err := users.Delete(user.Name, user.Password)
	if err != nil {
		log.Println(err)
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}

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
	}

	var msg = "success"
	if statusCode < 0 {
		msg = "failure"
	}
	c.JSON(http.StatusOK, gin.H{
		"statusMessage": msg,
	})
}
