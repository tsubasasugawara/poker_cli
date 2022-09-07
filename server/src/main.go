package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	router(engine)

    engine.Run(":8080")
}
