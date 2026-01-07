package middleware

import (
	"cg-tg-bot/utils"

	"github.com/gin-gonic/gin"
)

func CrossSite(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Key")
	c.Next()
}

func Auth(c *gin.Context) {
	authKey := c.GetHeader("key")

	if authKey != utils.AppConfig.AuthKey {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "不允许调用",
		})
		c.Abort()
		return
	}
}
