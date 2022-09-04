
package main

import (
	"net/http"

	"poker/controller/user"

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
		userEngine.POST("/register", user.Register)
		userEngine.POST("/login", user.Login)
		userEngine.POST("/logout", user.Logout)
		userEngine.POST("/delete", user.Delete)
		userEngine.POST("/edit", user.Edit)
	}

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"User-Agent": ua,
		})
	})

    engine.Run(":8080")
}
