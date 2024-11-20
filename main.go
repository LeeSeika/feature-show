package main

import (
	"gv_server/core"
	"gv_server/global"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitZap(&global.Config.ZapConfig)
	//初始化数据库连接
	global.Db = core.InitGorm()
	global.Log.Error("程序启动成功")
}
