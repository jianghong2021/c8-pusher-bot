package router

import (
	"cg-tg-bot/controller"
	"cg-tg-bot/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {

	app.GET("/", controller.Index)

	pusher := app.Group("/open/data/push", middleware.Auth)
	{
		pusher.POST("/payCashOut", controller.PushayCashOut)
		pusher.POST("/payRecharge", controller.PayRecharge)
		pusher.POST("/billiardEntertained", controller.BilliardEntertained)
		pusher.POST("/dailyFunds", controller.DailyFunds)
		pusher.POST("/lottery28Entertained", controller.Lottery28Entertained)
	}
}
