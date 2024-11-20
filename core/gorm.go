package core

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gv_server/global"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Error("mysql地址为空，取消连接。")
		return nil
	}
	dns := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.Mysql.LogLevel == "dev" {
		//开发模式显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: mysqlLogger},
	)
	if err != nil {
		global.Log.Error("数据库连接失败", zap.String("err", err.Error()))
		global.Log.Error("DNS:", zap.String("dns", dns))
		return nil
	}
	sqlDb, _ := db.DB()
	//最大连接数
	sqlDb.SetMaxIdleConns(10)
	//最多可容纳
	sqlDb.SetMaxOpenConns(100)
	//连接最大复用时间，不能超过mysql的wait_timeout
	sqlDb.SetConnMaxLifetime(time.Hour * 4)
	global.Log.Info("数据库连接成功")
	return db
}
