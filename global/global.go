package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gv_server/config"
)

var (
	Config = &config.Config{}
	Db     = &gorm.DB{}
	V      = &viper.Viper{}
)
