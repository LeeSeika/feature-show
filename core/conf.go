package core

import (
	"github.com/leeseika/feature-show/config"
	"github.com/leeseika/feature-show/global"
	"github.com/spf13/viper"
)

// InitCore 读取yaml配置文件
func InitCore() {
	//配置文件路径
	const ConfigFilePath = "settings.yaml"
	c := &config.Config{}
	v := viper.New()
	v.SetConfigFile(ConfigFilePath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(c)
	if err != nil {
		panic(err)
	}
	global.V = v
	global.Config = c
}
