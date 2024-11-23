package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/leeseika/feature-show/dao/mysql"
	"github.com/leeseika/feature-show/logger"
	"github.com/leeseika/feature-show/router"
	"github.com/leeseika/feature-show/services"
	"github.com/leeseika/feature-show/settings"
	"go.uber.org/zap"
)

func main() {
	var configFileName string
	flag.StringVar(&configFileName, "conf", "./settings.yaml", "settings file")

	err := settings.Init(configFileName)
	if err != nil {
		fmt.Printf("settings init error:%s\n", err)
		return
	}

	err = logger.Init()
	if err != nil {
		zap.L().Error("logger init error: ", zap.Error(err))
		return
	}
	zap.L().Info("successfully initialized logger")
	defer zap.L().Sync()

	db, err := mysql.Init()
	if err != nil {
		zap.L().Error("mysql init error: ", zap.Error(err))
		return
	}
	zap.L().Info("successfully connected to mysql")

	services.Init(db)

	engine := router.Setup()

	host := settings.Conf.App.Host
	port := settings.Conf.App.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()
	// graceful shutdown
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("failed to shutdown server: ", zap.Error(err))
	}
}
