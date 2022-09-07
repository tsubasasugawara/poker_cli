
package main

import (
	"net/http"
	"poker/controller/user"
	"poker/controller/rooms"

    "github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	var ua string
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	userEngine := engine.Group("/user")
	{
		userEngine.POST("/regist", user.Regist)
		userEngine.POST("/login", user.Login)
		userEngine.POST("/delete", user.Delete)
		userEngine.POST("/edit", user.Edit)
	}

	roomEngine := engine.Group("/room")
	{
		roomEngine.POST("/create", rooms.CreateRoom)
		roomEngine.POST("/participate", rooms.ParticipateRoom)
		roomEngine.POST("/exit", rooms.ExitRoom)
	}

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"User-Agent": ua,
		})
	})

    engine.Run(":8080")
}
