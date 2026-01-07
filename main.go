package main

import (
	"cg-tg-bot/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	router.InitRouter(app)

	app.Run(":8081")
}
