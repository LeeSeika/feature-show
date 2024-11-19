package config

type Config struct {
	Mysql Mysql `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
}
