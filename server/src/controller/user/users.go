package user

import (
	"net/http"

    "github.com/gin-gonic/gin"
)


func Regist(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})
}

func Login(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func Logout(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "logout",
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
