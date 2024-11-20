package routers

import (
	"github.com/gin-gonic/gin"
	"gv_server/global"
)

func InitRouter() *gin.Engine {
	//开启日志颜色
	gin.ForceConsoleColor()
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	return router
}
