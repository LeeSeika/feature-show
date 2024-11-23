package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeseika/feature-show/middleware"
	"github.com/leeseika/feature-show/settings"
)

func Setup() *gin.Engine {
	if settings.Conf.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "page not found",
		})
	})
	return engine
}
