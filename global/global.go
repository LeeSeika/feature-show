package global

import (
	"github.com/leeseika/feature-show/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config = &config.Config{}
	Db     = &gorm.DB{}
	V      = &viper.Viper{}
	Log    = &zap.Logger{}
)
