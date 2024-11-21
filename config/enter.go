package config

type Config struct {
	Mysql     Mysql     `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	ZapConfig ZapConfig `json:"zapConfig" yaml:"zap" mapstructure:"zap"`
	System    System    `json:"system" yaml:"system"`
}
