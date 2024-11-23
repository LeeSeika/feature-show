package settings

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(Config)

func Init(configFileName string) (err error) {
	viper.SetConfigFile(configFileName)
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config read:%s", err))
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic(fmt.Errorf("fatal error config unmarshal:%s", err))
	}

	return nil
}

type Config struct {
	*App   `mapstructure:"app"`
	*Log   `mapstructure:"log"`
	*MySQL `mapstructure:"mysql"`
}

type App struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	Level      string   `mapstructure:"level"`
	TimeFormat string   `mapstructure:"time_format"`
	Caller     bool     `mapstructure:"caller"`
	Stacktrace bool     `mapstructure:"stacktrace"`
	Encode     string   `mapstructure:"encode"`
	Writer     string   `mapstructure:"writer"`
	MaxAge     int      `mapstructure:"max_age"`
	MaxBackups int      `mapstructure:"max_backups"`
	MaxSize    int      `mapstructure:"max_size"`
	LogFile    *LogFile `mapstructure:"log_file"`
}

type LogFile struct {
	MaxSize  int      `mapstructure:"max_size"`
	MaxAge   int      `mapstructure:"max_age"`
	Backups  int      `mapstructure:"backups"`
	Compress bool     `mapstructure:"compress"`
	Output   []string `mapstructure:"output"`
}

type MySQL struct {
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DB              string `mapstructure:"db"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time"`
}
