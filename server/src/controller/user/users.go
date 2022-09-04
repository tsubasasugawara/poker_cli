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

	users := users.NewUsers()
	id, err := users.Regist(user.Name, user.Password)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Login(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func Delete(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "delete",
	})
}

func Edit(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "edit",
	})
}
