package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leeseika/feature-show/global"
)

func InitRouter() *gin.Engine {
	//开启日志颜色
	gin.ForceConsoleColor()
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 示例接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
