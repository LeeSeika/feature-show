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
	return router
}
