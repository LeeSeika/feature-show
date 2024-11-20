package config

import "strconv"

type Mysql struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Db       string `json:"db" yaml:"db" mapstructure:"db"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	LogLevel string `json:"log_level" yaml:"log_level" mapstructure:"log_level"` //日志等级，debug就是输出全部sql，
}

func (m *Mysql) Dsn() string {
	return m.User + ":" +
		m.Password +
		"@tcp(" +
		m.Host +
		":" +
		strconv.Itoa(m.Port) + ")/" +
		m.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
}
