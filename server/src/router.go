package main

import (
	"poker/controller/play"
	"poker/controller/user"
	"poker/controller/rooms"

    "github.com/gin-gonic/gin"
)

func router(engine *gin.Engine) {
	hub := play.NewHub()
	go hub.Run()
	engine.GET("/play", func(c *gin.Context) {
		play.ServeWs(hub, c)
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
}
