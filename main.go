package main

import (
	"github.com/leeseika/feature-show/core"
	"github.com/leeseika/feature-show/global"
	"github.com/leeseika/feature-show/routers"
	"go.uber.org/zap"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitZap(&global.Config.ZapConfig)
	//初始化数据库连接
	global.Db = core.InitGorm()
	//初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.GetAddr()
	global.Log.Info("程序运行端口地址:", zap.String("addr:", addr))
	router.Run(addr)
}
